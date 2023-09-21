package textsplitter

import (
	"github.com/aresa7796/sentences"
	"github.com/aresa7796/sentences/english"
	"github.com/pkoukk/tiktoken-go"
	"regexp"
	"strings"
)

const CHUNKING_REGEX = "[^,.;。]+[,.;。]?"
const DEFUALT_PARAGRAPH_SEP = "\n\n\n"
const separator = " "

const TokenEncoding = "cl100k_base"

var AllowedSpecial = []string{}
var DisallowedSpecial = []string{"all"}

// Split represents a text split.
type Split struct {
	Text       string
	IsSentence bool
}

// SentenceSplitter splits text into chunks with a preference for complete sentences.
type SentenceSplitter struct {
	ChunkSize              int
	ChunkOverlap           int
	Separator              string
	ParagraphSeparator     string
	SecondaryChunkingRegex string
	chunkingTokenizerFn    *sentences.DefaultSentenceTokenizer
	Tokenizer              *tiktoken.Tiktoken
}

// NewSentenceSplitter creates a new SentenceSplitter instance.
func NewSentenceSplitter(chunkSize, chunkOverlap int) *SentenceSplitter {
	tokenizer, _ := english.NewSentenceTokenizer(nil)
	tk, _ := tiktoken.GetEncoding(TokenEncoding)
	return &SentenceSplitter{
		ChunkSize:              chunkSize,
		ChunkOverlap:           chunkOverlap,
		Separator:              separator,
		ParagraphSeparator:     DEFUALT_PARAGRAPH_SEP,
		SecondaryChunkingRegex: CHUNKING_REGEX,
		chunkingTokenizerFn:    tokenizer,
		Tokenizer:              tk,
	}
}

// SplitTextMetadataAware splits text with metadata into chunks.
func (s *SentenceSplitter) SplitTextMetadataAware(text, metadataStr string) []string {
	metadataLen := len(s.TokenEncode(metadataStr))
	//metadataLen := utf8.RuneCountInString(metadataStr)
	effectiveChunkSize := s.ChunkSize - metadataLen
	return s.splitText(text, effectiveChunkSize)
}

// SplitText splits text into chunks.
func (s *SentenceSplitter) SplitText(text string) []string {
	return s.splitText(text, s.ChunkSize)
}

func (s *SentenceSplitter) splitText(text string, chunkSize int) []string {
	if text == "" {
		return nil
	}

	splits := s.split(text, chunkSize)
	chunks := s.merge(splits, chunkSize)

	return chunks
}

func (s *SentenceSplitter) split(text string, chunkSize int) []Split {
	if len(s.TokenEncode(text)) <= chunkSize {
		//if utf8.RuneCountInString(text) <= chunkSize {
		return []Split{{Text: text, IsSentence: true}}
	}

	var splits []string
	for _, splitFn := range []func(string) []string{
		s.SplitBySep(s.ParagraphSeparator, true),
		s.ChunkingTokenizerFn(),
	} {
		splits = splitFn(text)
		if len(splits) > 1 {
			break
		}
	}

	var isSentence bool
	if len(splits) > 1 {
		isSentence = true
	} else {
		for _, splitFn := range []func(string) []string{
			s.SplitByRegex(s.SecondaryChunkingRegex),
			s.SplitBySep(s.Separator, true),
			s.SplitByChar(),
		} {
			splits = splitFn(text)
			if len(splits) > 1 {
				break
			}
		}
		isSentence = false
	}

	var newSplits []Split
	for _, split := range splits {
		splitLen := len(s.TokenEncode(split))
		//splitLen := utf8.RuneCountInString(split)
		if splitLen <= chunkSize {
			newSplits = append(newSplits, Split{Text: split, IsSentence: isSentence})
		} else {
			ns := s.split(split, chunkSize)
			if len(ns) == 0 {
				// Handle 0 length split
				continue
			}
			newSplits = append(newSplits, ns...)
		}
	}

	return newSplits
}

func (s *SentenceSplitter) merge(splits []Split, chunkSize int) []string {
	var chunks []string
	var curChunk []string
	curChunkLen := 0

	for len(splits) > 0 {
		curSplit := splits[0]
		curSplitLen := len(s.TokenEncode(curSplit.Text))
		//curSplitLen := utf8.RuneCountInString(curSplit.Text)
		if curSplitLen > chunkSize {
			panic("Single token exceeds chunk size")
		}
		if curChunkLen+curSplitLen > chunkSize && len(curChunk) > 0 {
			// If adding split to current chunk exceeds chunk size, close out chunk
			chunks = append(chunks, strings.Join(curChunk, ""))
			curChunk = nil
			curChunkLen = 0
		} else {
			if curSplit.IsSentence || curChunkLen+curSplitLen < chunkSize-s.ChunkOverlap || len(curChunk) == 0 {
				// Add split to chunk
				curChunkLen += curSplitLen
				curChunk = append(curChunk, curSplit.Text)
				splits = splits[1:]
			} else {
				// Close out chunk
				chunks = append(chunks, strings.Join(curChunk, ""))
				curChunk = nil
				curChunkLen = 0
			}
		}
	}

	// Handle the last chunk
	chunk := strings.Join(curChunk, "")
	if chunk != "" {
		chunks = append(chunks, chunk)
	}

	// Run postprocessing to remove blank spaces
	chunks = s.postprocessChunks(chunks)

	return chunks
}

func (s *SentenceSplitter) postprocessChunks(chunks []string) []string {
	var newChunks []string
	for _, doc := range chunks {
		if strings.ReplaceAll(doc, " ", "") == "" {
			continue
		}
		newChunks = append(newChunks, doc)
	}
	return newChunks
}

func (s *SentenceSplitter) TokenEncode(text string) []int {
	return s.Tokenizer.Encode(text, AllowedSpecial, DisallowedSpecial)
}

// SplitByChar splits text by character.
func (s *SentenceSplitter) SplitByChar() func(string) []string {
	return func(text string) []string {
		result := make([]string, len(text))
		for i, char := range text {
			result[i] = string(char)
		}
		return result
	}
}

// SplitByRegex splits text by regex.
func (s *SentenceSplitter) SplitByRegex(regex string) func(string) []string {
	re := regexp.MustCompile(regex)
	return func(text string) []string {
		return re.FindAllString(text, -1)
	}
}

func (s *SentenceSplitter) ChunkingTokenizerFn() func(string) []string {
	return func(text string) []string {
		return s.chunkingTokenizerFn.Tokenize2String(text)
	}
}

// SplitBySep splits text by separator.
func (s *SentenceSplitter) SplitBySep(sep string, keepSep bool) func(string) []string {
	if keepSep {
		return func(text string) []string {
			return splitTextKeepSeparator(text, sep)
		}
	} else {
		return func(text string) []string {
			return strings.Split(text, sep)
		}
	}
}

// splitTextKeepSeparator is a helper function to split text while keeping the separator.
func splitTextKeepSeparator(text, sep string) []string {
	var result []string
	splitText := strings.Split(text, sep)
	for i, s := range splitText {
		if i < len(splitText)-1 {
			result = append(result, s+sep)
		} else {
			result = append(result, s)
		}
	}
	return result
}

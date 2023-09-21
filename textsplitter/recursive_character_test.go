package textsplitter

import (
	"fmt"
	"testing"
	"unicode/utf8"

	"github.com/aresa7796/langchaingo/schema"
	"github.com/stretchr/testify/assert"
)

//nolint:dupword
func TestRecursiveCharacterSplitter(t *testing.T) {
	t.Parallel()
	type testCase struct {
		text         string
		chunkOverlap int
		chunkSize    int
		expectedDocs []schema.Document
	}
	testCases := []testCase{
		{
			text:         "Hi.\nI'm Harrison.\n\nHow?\na\nb",
			chunkOverlap: 1,
			chunkSize:    20,
			expectedDocs: []schema.Document{
				{PageContent: "Hi.\nI'm Harrison.", Metadata: map[string]any{}},
				{PageContent: "How?\na\nb", Metadata: map[string]any{}},
			},
		},
		{
			text:         "Hi.\nI'm Harrison.\n\nHow?\na\nbHi.\nI'm Harrison.\n\nHow?\na\nb",
			chunkOverlap: 1,
			chunkSize:    40,
			expectedDocs: []schema.Document{
				{PageContent: "Hi.\nI'm Harrison.", Metadata: map[string]any{}},
				{PageContent: "How?\na\nbHi.\nI'm Harrison.\n\nHow?\na\nb", Metadata: map[string]any{}},
			},
		},
		{
			text:         "name: Harrison\nage: 30",
			chunkOverlap: 1,
			chunkSize:    40,
			expectedDocs: []schema.Document{
				{PageContent: "name: Harrison\nage: 30", Metadata: map[string]any{}},
			},
		},
		{
			text: `name: Harrison
age: 30

name: Joe
age: 32`,
			chunkOverlap: 1,
			chunkSize:    40,
			expectedDocs: []schema.Document{
				{PageContent: "name: Harrison\nage: 30", Metadata: map[string]any{}},
				{PageContent: "name: Joe\nage: 32", Metadata: map[string]any{}},
			},
		},
		{
			text: `Hi.
I'm Harrison.

How? Are? You?
Okay then f f f f.
This is a weird text to write, but gotta test the splittingggg some how.

Bye!

-H.`,
			chunkOverlap: 1,
			chunkSize:    10,
			expectedDocs: []schema.Document{
				{PageContent: "Hi.", Metadata: map[string]any{}},
				{PageContent: "I'm", Metadata: map[string]any{}},
				{PageContent: "Harrison.", Metadata: map[string]any{}},
				{PageContent: "How? Are?", Metadata: map[string]any{}},
				{PageContent: "You?", Metadata: map[string]any{}},
				{PageContent: "Okay then", Metadata: map[string]any{}},
				{PageContent: "f f f f.", Metadata: map[string]any{}},
				{PageContent: "This is a", Metadata: map[string]any{}},
				{PageContent: "a weird", Metadata: map[string]any{}},
				{PageContent: "text to", Metadata: map[string]any{}},
				{PageContent: "write, but", Metadata: map[string]any{}},
				{PageContent: "gotta test", Metadata: map[string]any{}},
				{PageContent: "the", Metadata: map[string]any{}},
				{PageContent: "splittingg", Metadata: map[string]any{}},
				{PageContent: "ggg", Metadata: map[string]any{}},
				{PageContent: "some how.", Metadata: map[string]any{}},
				{PageContent: "Bye!\n\n-H.", Metadata: map[string]any{}},
			},
		},
	}
	splitter := NewRecursiveCharacter()
	for _, tc := range testCases {
		splitter.ChunkOverlap = tc.chunkOverlap
		splitter.ChunkSize = tc.chunkSize

		docs, err := CreateDocuments(splitter, []string{tc.text}, nil)
		assert.NoError(t, err)
		assert.Equal(t, tc.expectedDocs, docs)
	}
}

func TestRecursiveCharacter_SplitText(t *testing.T) {
	splitter := NewRecursiveCharacter()
	splitter.ChunkOverlap = 0
	splitter.ChunkSize = 400

	text := `738
00:21:47,240 --> 00:21:48,240
有可能你這個

739
00:21:48,240 --> 00:21:50,240
也是進到了腹膜前間性

740
00:21:50,240 --> 00:21:51,240
或者是在腸管裡面

741
00:21:51,240 --> 00:21:53,240
因為這個翹是很脆的

742
00:21:53,240 --> 00:21:54,240
稍微前面阻力大了以後

743
00:21:54,240 --> 00:21:56,240
它就會彎曲

744
00:21:56,240 --> 00:21:57,240
那麼這個是這個

745
00:21:57,240 --> 00:21:59,240
穿刺翹的一個優點

746
00:21:59,240 --> 00:22:00,240
那麼這個優點呢

747
00:22:00,240 --> 00:22:01,240
同時也有一個缺點

748
00:22:01,240 --> 00:22:02,240
什麼缺點呢

749
00:22:02,240 --> 00:22:04,240
因為這個針是比較細的

750
00:22:04,240 --> 00:22:05,240
也比較銳

751
00:22:05,240 --> 00:22:06,240
所以在穿刺過程當中

752
00:22:06,240 --> 00:22:07,240
這種突破感呢

753
00:22:07,240 --> 00:22:09,240
需要仔細體會

754
00:22:09,240 --> 00:22:10,240
那麼對於一些經驗豐富的

755
00:22:10,240 --> 00:22:11,240
大夫沒有問題

756
00:22:11,240 --> 00:22:12,240
那麼對於一些

757
00:22:12,240 --> 00:22:13,240
那個新手呢

758
00:22:13,240 --> 00:22:14,240
可能這個突破感

759
00:22:14,240 --> 00:22:15,240
會稍微弱一點

760
00:22:15,240 --> 00:22:16,240
當然如果是在

761
00:22:16,240 --> 00:22:17,240
B腸輪導下呢

762
00:22:17,240 --> 00:22:18,240
這也沒有問題

763
00:22:18,240 --> 00:22:19,240
對於這個

764
00:22:19,240 --> 00:22:21,240
這個紅色的這個針呢

765
00:22:21,240 --> 00:22:22,240
穿刺掏包裡面這個針

766
00:22:22,240 --> 00:22:23,240
相對粗一點

767
00:22:23,240 --> 00:22:24,240
那麼它的這個針的這個

768
00:22:24,240 --> 00:22:25,240
穿刺呢

769
00:22:25,240 --> 00:22:27,240
相對突破感更強一些

770
00:22:27,240 --> 00:22:28,240
那麼當然它也挺銳利的

771
00:22:28,240 --> 00:22:29,240
那麼穿刺過程當中

772
00:22:29,240 --> 00:22:30,240
要注意損傷

773
00:22:30,240 --> 00:22:31,240
那麼我們如果是

774
00:22:31,240 --> 00:22:32,240
突破感很強

775
00:22:32,240 --> 00:22:34,240
如果在B腸輪導下呢

776
00:22:34,240 --> 00:22:35,240
也會看得比較清楚

777
00:22:35,240 --> 00:22:36,240
針相對粗一點

778
00:22:36,240 --> 00:22:37,240
那麼現在呢

779
00:22:37,240 --> 00:22:38,240
這兩種穿刺針

780
00:22:38,240 --> 00:22:39,240
我們都在用

781
00:22:39,240 --> 00:22:40,240
第三種呢

782
00:22:40,240 --> 00:22:41,240
就是這個氣止針

783
00:22:41,240 --> 00:22:42,240
那麼氣止針

784
00:22:42,240 --> 00:22:43,240
我們也很喜歡

785
00:22:43,240 --> 00:22:44,240
而且用得也比較齡暢

786
00:22:44,240 --> 00:22:45,240
那麼我們在做腹腔鏡的時候

787
00:22:45,240 --> 00:22:46,240
用的是這個針

788
00:22:46,240 --> 00:22:47,240
大家看這個針的

789
00:22:47,240 --> 00:22:48,240
這個針頭是凸的

790
00:22:48,240 --> 00:22:49,240
那麼它的這個針尖

791
00:22:49,240 --> 00:22:51,240
是在這個後面藏著的

792
00:22:51,240 --> 00:22:52,240
那麼有阻力的時候

793
00:22:52,240 --> 00:22:53,240
這個是能進去的

794
00:22:53,240 --> 00:22:54,240
但是呢要提醒的是

795
00:22:54,240 --> 00:22:56,240
這個氣止針的話

796
00:22:56,240 --> 00:22:57,240
會有一個問題

797
00:22:57,240 --> 00:22:58,240
什麼問題呢

798
00:22:58,240 --> 00:22:59,240
就是如果這個病人

799
00:22:59,240 --> 00:23:01,240
有腹膜前間隙的過程當中

800
00:23:01,240 --> 00:23:02,240
那麼很有可能會
`
	docs, err := CreateDocuments(splitter, []string{text}, nil)
	assert.NoError(t, err)
	for i, doc := range docs {
		fmt.Printf("Chunk [%d] : %s\n", i, doc.PageContent)
	}

}

func TestLength(t *testing.T) {
	text := `738
00:21:47,240 --> 00:21:48,240
有可能你這個

739
00:21:48,240 --> 00:21:50,240
也是進到了腹膜前間性

740
00:21:50,240 --> 00:21:51,240
或者是在腸管裡面

741
00:21:51,240 --> 00:21:53,240
因為這個翹是很脆的

742
00:21:53,240 --> 00:21:54,240
稍微前面阻力大了以後
hello
743
00:21:54,240 --> 00:21:56,240
它就會彎曲`
	fmt.Println(utf8.RuneCountInString(text))
}

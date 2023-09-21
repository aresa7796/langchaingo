package textsplitter

import (
	"fmt"
	"testing"

	"github.com/aresa7796/langchaingo/schema"
	"github.com/stretchr/testify/assert"
)

func TestTokenSplitter(t *testing.T) {
	t.Parallel()
	type testCase struct {
		text         string
		chunkOverlap int
		chunkSize    int
		expectedDocs []schema.Document
	}
	//nolint:dupword
	testCases := []testCase{
		{
			text:         "Hi.\nI'm Harrison.\n\nHow?\na\nb",
			chunkOverlap: 1,
			chunkSize:    20,
			expectedDocs: []schema.Document{
				{PageContent: "Hi.\nI'm Harrison.\n\nHow?\na\nb", Metadata: map[string]any{}},
			},
		},
		{
			text:         "Hi.\nI'm Harrison.\n\nHow?\na\nbHi.\nI'm Harrison.\n\nHow?\na\nb",
			chunkOverlap: 1,
			chunkSize:    40,
			expectedDocs: []schema.Document{
				{PageContent: "Hi.\nI'm Harrison.\n\nHow?\na\nbHi.\nI'm Harrison.\n\nHow?\na\nb", Metadata: map[string]any{}},
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
				{PageContent: "name: Harrison\nage: 30\n\nname: Joe\nage: 32", Metadata: map[string]any{}},
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
				{PageContent: "Hi.\nI'm Harrison.\n\nHow? Are?", Metadata: map[string]any{}},
				{PageContent: "? You?\nOkay then f f f f.\n", Metadata: map[string]any{}},
				{PageContent: ".\nThis is a weird text to write, but", Metadata: map[string]any{}},
				{PageContent: " but gotta test the splittingggg some how.\n\n", Metadata: map[string]any{}},
				{PageContent: ".\n\nBye!\n\n-H.", Metadata: map[string]any{}},
			},
		},
	}
	splitter := NewTokenSplitter()
	for _, tc := range testCases {
		splitter.ChunkOverlap = tc.chunkOverlap
		splitter.ChunkSize = tc.chunkSize

		docs, err := CreateDocuments(splitter, []string{tc.text}, nil)
		assert.NoError(t, err)
		assert.Equal(t, tc.expectedDocs, docs)
	}
}

func TestToken_SplitText(t *testing.T) {
	splitter := NewTokenSplitter()
	splitter.ChunkOverlap = 30
	splitter.ChunkSize = 400

	text := `2018 Psych | 抑郁的严重程度与治疗：对两个争议性问题进行综述
2019-04-11 灵北学院
66
在佛罗里达州奥兰多市举行的2018年心理学大会开幕式全体会议上，来自罗德岛普罗维登斯布朗大学精神病学和人类行为学教授Mark Zimmerman发表了讲话，他指出抗抑郁药物不仅限于重度抑郁患者，心理疗法也不仅限于轻度和中度抑郁患者。他补充说，由于在评估抑郁患者的心理治疗和抗抑郁药物治疗的临床试验中排除了个别患者，因此目前给出明确建议的能力是非常有限的；此外对于轻度、中度和重度抑郁的划分标准也存在差异。




Zimmerman教授指出，抑郁的严重程度会影响治疗决策和治疗指南建议（例如，美国精神病学协会(APA)和英国国家健康和护理卓越研究所(NICE)），但轻度、中度和重度抑郁之间的界限既不明确也不一致。他强调，在对抑郁严重程度进行概念化时没有明确的共识，文献中也没有关于临床医生如何划分严重程度的数据。


|对抑郁严重程度进行概念化时没有明确的共识


根据第5版精神障碍诊断和统计手册(DSM-5)，轻度、中度和重度抑郁之间的区别是基于症状数量、症状强度、抑郁程度、可管控性以及其对社会和职业功能的影响等方面。


相反：

·  关于抑郁治疗的研究均使用抑郁评定量表来区分轻度、中度和重度抑郁，但目前存在多种量表，且各自的内容和评级说明存在差异——因此，这些量表对抑郁严重程度的分类也各不相同，使用不同的量表对患者的抑郁严重程度进行判断，所得出的结论也是不同的


·  目前尚不清楚临床医生如何评定抑郁的严重程度，所以在全体会议期间，Zimmerman教授要求所有临床医生完成调查问卷，评估他们在评定一个患者是否患有重度抑郁时所考虑的因素，并根据重要性进行排序——这些因素包括症状的数量、强度、频率、顽固性、功能受损程度、症状持续时间、难以应对压力、生活满意度以及生活质量下降


|需要使用可靠、有效和临床上有用的方法来评定抑郁的严重程度


Zimmerman教授强调，有必要采用可靠、有效和临床上有用的方法来评定抑郁的严重程度。他介绍了他之前进行的一项研究，该研究中245名门诊抑郁患者完成了三项自我报告量表，即临床实用抑郁转归量表(CUDOS)、抑郁症状快速自评量表(QIDS)和患者健康问卷抑郁筛查量表(PHQ­9)。此外，也使用了汉密尔顿抑郁量表（17项）对患者的抑郁严重程度进行评定。


·  当使用临床实用抑郁转归量表和汉密尔顿抑郁量表（17项）时，中度抑郁是最常见的严重程度类别


·  当使用患者健康问卷抑郁筛查量表和抑郁症状快速自评量表时，大多数患者被归类为重度抑郁


Zimmerman教授指出，与使用患者健康问卷抑郁筛查量表和抑郁症状快速自评量表相比，使用临床实用抑郁转归量表对患者进行评定时，被归为重度抑郁患者的人数明显减少。




Zimmerman教授指出，抗抑郁药物是所有药物处方中最常用的药物之一，但它们的有效性一直存在争议。


因为不同的研究者使用不同的临界值来界定重度抑郁。例如，一些研究者认为当HAM-D≥20时为重度抑郁，而其他研究者则认为当HAM-D≥28时才可评定为重度抑郁。


|不同的研究中使用了不同的临界值来定义重度抑郁


Zimmerman教授补充说，官方治疗指南会根据抑郁的严重程度来提供建议。 美国精神病学协会建议：


·  针对轻度和中度抑郁采用药物疗法或心理疗法

·  针对重度抑郁采用药物疗法


英国国家健康和护理卓越研究所指南建议：


·  针对轻度抑郁采用心理疗法

·  针对中度和重度抑郁采用药物疗法和心理疗法


|需要提高可靠性和有效性


美国食品和药物管理局(FDA)数据库对45项临床试验进行分析后发现，在试验中用于重度抑郁患者的抗抑郁药物和安慰剂之间存在较多统计学差异：


·  就使用抗抑郁药物的治疗小组来说，症状减轻状况与汉密尔顿抑郁量表的平均初始评分显著相关，评分越高，变化越大

·  就使用安慰剂的治疗小组来说，汉密尔顿抑郁量表的平均初始评分越高，变化越小


早期停药在汉密尔顿抑郁量表平均初始评分较高的患者中更常见。研究得出结论，这些数据可能为未来抗抑郁药临床试验设计提供相关信息。然而，Zimmerman教授指出，这些研究结果并不适用于临床实践，由于严格的纳入和排除标准，上述患者中的多数人并不符合参加临床试验的资格。


|初始严重程度与抗抑郁药物疗效之间的关系可归因于重度抑郁患者对安慰剂的反应性降低，而不是对药物的反应性增强了


2008年，美国食品与药品管理局的数据库对35项抗抑郁药物临床试验的进一步分析后得出结论，初始严重程度与抗抑郁药物疗效之间的关系可归因于重度抑郁患者对安慰剂的反应性降低，而不是对药物的反应性增强了。这一结论， 是根据英国国家健康和护理卓越研究所提供的指南并基于汉密尔顿抑郁量表的3点变化而得出的，引发了关于抗抑郁药物是否有效的激烈争论。


Zimmerman教授解释说，这一研究有一定的局限性，因为：


·  研究使用平均值，而不是患者的各自分数，导致分数分布太过分散，同时也未能准确识别轻度、中度和重度抑郁患者

·  研究者对病症缓解的平均变化进行分析，而不是病症缓解患者所占的的百分比

·  对样本患者抑郁严重程度的识别并不准确，重度抑郁可能被归为轻度抑郁

·  普适性比较有限，因为许多患者并不符合药效试验的标准，主要是因为他们的症状不够严重




国家心理健康研究所合作研究调查了两种简短的心理疗法，即人际心理治疗和认知行为治疗(CBT)，主要用于门诊抑郁患者的治疗。Zimmerman教授讲述，250名患者被随机分配到以下几个小组：


·  人际关系心理治疗小组

·  认知行为治疗小组

·  抗抑郁药物加临床管理小组（作为标准对照治疗）

·  安慰剂加临床管理小组


|抗抑郁药物加临床管理疗效最为明显，安慰剂加临床管理疗效最差


所有治疗小组的患者均表现出抑郁症状的明显缓解和功能改善。抗抑郁药物加临床管理疗效最为明显，安慰剂加临床管理疗效最差。关于心理疗法对重度抑郁无效的研究建议具有争议性，并且正在被重新研究。



参考文献

1.    Zimmerman M, et al. J Clin Psych 2012

2.    DeRubeis R, et al. Am J Psych 1999

3.    Gibbons R, et al. Arch Gen Psychiatry. 2012

4.    Khan A, et al. J Clin Psychopharmacol 2002

5.    American Psychiatric Association Practice Guideline for the Treatment of Patients With Major Depressive Disorder, Third Edition, 2010. . Accessed 25 October 2018.

6.    National Collaborating Centre for Mental Health (UK) NICE Clinical Guideline no 90. Depression: The Treatment and Management of Depression in Adults (Updated Edition).. Accessed 25 October 2018.

7.    Kirsch I. PLoS Med 2008

8.    Elkin I,  et al. Arch Gen Psych 1989

`
	docs, err := CreateDocuments(splitter, []string{text}, nil)
	assert.NoError(t, err)
	for i, doc := range docs {
		fmt.Printf("Chunk [%d] : %s\n", i, doc.PageContent)
	}

}

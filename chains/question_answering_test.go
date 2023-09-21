package chains

import (
	"context"
	"github.com/aresa7796/langchaingo/prompts"
	"github.com/aresa7796/langchaingo/schema"
	"os"
	"strings"
	"testing"

	"github.com/aresa7796/langchaingo/llms/openai"
	"github.com/stretchr/testify/require"
)

func TestRefineQA(t *testing.T) {
	t.Parallel()

	//if openaiKey := os.Getenv("OPENAI_API_KEY"); openaiKey == "" {
	//	t.Skip("OPENAI_API_KEY not set")
	//}
	openaiBaseURLOption := openai.WithBaseURL("")
	openaiApiKeyOption := openai.WithToken("")
	openAIModelOption := openai.WithModel("gpt-3.5-turbo-0613")
	llm, err := openai.New(openaiBaseURLOption, openaiApiKeyOption, openAIModelOption)
	require.NoError(t, err)

	docs := loadTestData(t)

	qaChain := LoadRefineQA(llm)

	results, err := Call(
		context.Background(),
		qaChain,
		map[string]any{
			"input_documents": docs,
			"question":        "What is the name of the lion?",
		},
	)
	require.NoError(t, err)

	_, ok := results["text"].(string)
	t.Log(results["text"])
	require.True(t, ok, "result does not contain text key")
}

func TestMapReduceQA(t *testing.T) {
	t.Parallel()

	if openaiKey := os.Getenv("OPENAI_API_KEY"); openaiKey == "" {
		t.Skip("OPENAI_API_KEY not set")
	}
	llm, err := openai.New()
	require.NoError(t, err)

	docs := loadTestData(t)
	qaChain := LoadMapReduceQA(llm)

	result, err := Predict(
		context.Background(),
		qaChain,
		map[string]any{
			"input_documents": docs,
			"question":        "What is the name of the lion?",
		},
	)

	require.NoError(t, err)
	require.True(t, strings.Contains(result, "Leo"), "result does not contain correct answer Leo")
}

func TestMapRerankQA(t *testing.T) {
	t.Parallel()

	if openaiKey := os.Getenv("OPENAI_API_KEY"); openaiKey == "" {
		t.Skip("OPENAI_API_KEY not set")
	}
	llm, err := openai.New()
	require.NoError(t, err)

	docs := loadTestData(t)
	mapRerankChain := LoadMapRerankQA(llm)

	results, err := Call(
		context.Background(),
		mapRerankChain,
		map[string]any{
			"input_documents": docs,
			"question":        "What is the name of the lion?",
		},
	)

	require.NoError(t, err)

	answer, ok := results["text"].(string)
	require.True(t, ok, "result does not contain text key")
	require.True(t, strings.Contains(answer, "Leo"), "result does not contain correct answer Leo")
}

func TestRefineQA_local(t *testing.T) {
	t.Parallel()

	//if openaiKey := os.Getenv("OPENAI_API_KEY"); openaiKey == "" {
	//	t.Skip("OPENAI_API_KEY not set")
	//}
	openaiBaseURLOption := openai.WithBaseURL("")
	openaiApiKeyOption := openai.WithToken("")
	openAIModelOption := openai.WithModel("gpt-3.5-turbo-0613")
	llm, err := openai.New(openaiBaseURLOption, openaiApiKeyOption, openAIModelOption)
	require.NoError(t, err)

	//docs := loadTestData(t)
	docs := []schema.Document{}
	docs = append(docs, schema.Document{
		PageContent: "Source [1] /api/files/download?id=FwY86RDt7dLAlXoI&preview=1\n么下一步的,也就是我们,为什么要成立这个学院,我们的思考,我们将来的展望,我们到底是,PD优先,那我们是不是,都要做APD,那么APD的机器,怎么来,这管路怎么来,所以我想今天呢,我的这个时间呢,也不多,我也是希望,大家呢,进一步讨论,还有就是APD的定价,现在的收费很低,它是不是合理,怎么体现,我们医务人员的,劳动价值等等,那我们能看到,从我们整个,中国的这个数据来看,我们每年新发的尿毒症,大概是2000到3000,我们刚才登记的数据,大家那么努力,登记下来,这几年下来,也不过500例左右,但是我们全国,每年新发的,就有2000到3000人,所以呢,这个是我们的,想到的压力和挑战,所以我们要,多做这方面的工作,那么APD呢,它是有很好的优势,我觉得我们要开展的话,我们能够更多的,从APD的角度,来进行这方面的工作,那我们前面也做了一些,铺垫和基础,2015年呢,我们\n这个生命源泉的,这个基金呢,就有向全国各地的,这个推广的,这么一个工作,大家呢,也非常的积极和努力,我们现有的这些单位呢,都已经在开展,我们有一定的,这个捐赠的量,但是还很少,所以我觉得一定要,从这个APD机器上面,能够更多的,推动这个工作,那我们知道,重庆儿童医院呢,他们又有匹配,又自己也成立了,这样一个基金,所以我觉得,这个呢是,希望大家呢,从这方面呢,更多的一些思考,那么今天呢,我们这个猪胜之,中国儿童尿毒症的,推广项目呢,就是另外一块,所以我觉得,从这个角度呢,我们是可以去,争取一些资源,给这个APD机器呢,给到大家,能够进一步,推动这个工作,我们大家知道,这个是我们的,这个儿科的,这个前辈,先驱,他是我们国家的,37年的,首任的,中华儿科学会的会长,那么曾经是,担任我们医学院的,儿科教授,",
	}, schema.Document{
		PageContent: "Source [2] /api/files/download?id=woVdpzoaYq891S1l&preview=1\n么下一步的,也就是我们,为什么要成立这个学院,我们的思考,我们将来的展望,我们到底是,PD优先,那我们是不是,都要做APD,那么APD的机器,怎么来,这管路怎么来,所以我想今天呢,我的这个时间呢,也不多,我也是希望,大家呢,进一步讨论,还有就是APD的定价,现在的收费很低,它是不是合理,怎么体现,我们医务人员的,劳动价值等等,那我们能看到,从我们整个,中国的这个数据来看,我们每年新发的尿毒症,大概是2000到3000,我们刚才登记的数据,大家那么努力,登记下来,这几年下来,也不过500例左右,但是我们全国,每年新发的,就有2000到3000人,所以呢,这个是我们的,想到的压力和挑战,所以我们要,多做这方面的工作,那么APD呢,它是有很好的优势,我觉得我们要开展的话,我们能够更多的,从APD的角度,来进行这方面的工作,那我们前面也做了一些,铺垫和基础,2015年呢,我们\n这个生命源泉的,这个基金呢,就有向全国各地的,这个推广的,这么一个工作,大家呢,也非常的积极和努力,我们现有的这些单位呢,都已经在开展,我们有一定的,这个捐赠的量,但是还很少,所以我觉得一定要,从这个APD机器上面,能够更多的,推动这个工作,那我们知道,重庆儿童医院呢,他们又有匹配,又自己也成立了,这样一个基金,所以我觉得,这个呢是,希望大家呢,从这方面呢,更多的一些思考,那么今天呢,我们这个猪胜之,中国儿童尿毒症的,推广项目呢,就是另外一块,所以我觉得,从这个角度呢,我们是可以去,争取一些资源,给这个APD机器呢,给到大家,能够进一步,推动这个工作,我们大家知道,这个是我们的,这个儿科的,这个前辈,先驱,他是我们国家的,37年的,首任的,中华儿科学会的会长,那么曾经是,担任我们医学院的,儿科教授,",
	}, schema.Document{
		PageContent: "Source [3] \n的12个月死亡审查技术生存率分别为91.6%和93.5%（图1），两组的技术生存率无显著性差异（log-rank=0.018，p=0.894）。\n;\n图1\n无腹膜炎或无菌血症生存率：APD和HD的12个月无腹膜炎或无菌血症生存率分别为94.4%和93.6%（图2），两组的无腹膜炎或无菌血症生存率无显著性差异（log-rank=0.159，p=0.690）。\n;\n图2\n患者生存率：APD和HD的12个月患者生存率分别为96.7%和96.9%（图3），两组的患者生存率无显著性差异（log-rank=0.003，p=0.957）。\n;\n图3\n**0** **4**\n结论\n在紧急起始透析的ESRD患者中，与使用CVC进行HD相比，使用APD作为紧急起始透析的方式，可减少短期透析相关并发症且费用更低。\n;\n**0** **5**\n推荐理由\n从PD从业者发文数量上看，人们对紧急起始腹膜透析的兴趣\n增加。笔者进行前瞻性队列研究，比较行APD和HD进行紧急起始透析的ESRD患者的临床结局，结果表明，在紧急起始透析的ESRD患者中，与使用CVC进行HD相比，使用APD作为紧急起始透析的方式，可减少短期透析相关并发症且费用更低。\n以上文献出处：\nInt J Artif Organs. 2022 Aug;;45(8):;672-679.;\n;\n**END**\n;\nCN/MA/22/01-1103-01\n预览时标签不可点\n​\n**写留言**\n取消\n留言\n**我的留言**\n写留言\n展开我的留言\n留言被精选后将公开\n**精选留言**\n写留言\n写留言\n全部留言\n已无更多数据\n发消息\n关闭\n**写留言**\n提交更多\n正在加载\n表情\n正在加载\n**留言**\n微信扫一扫\n关注该公众号\n知道了\n微信扫一扫\n使用小程序\n取消允许\n取消允许\n：，。;视频小程序赞，轻点两下取消赞在看，轻点两下取消在看",
	})

	//qaChain := LoadRefineQA(llm)

	questionPrompt := prompts.NewPromptTemplate(
		`Use the following pieces of context to answer the question at the end. If you don't know the answer, just say that you don't know, don't try to make up an answer.

{{.context}}

Question: {{.question}}
INSTRUCTIONS:
- The format you return: Make sure to cite results using [number] notation after the reference. But don't list the references.
Helpful Answer:`,
		[]string{"context", "question"},
	)
	refinePrompt := prompts.NewPromptTemplate(
		`The original question is as follows: {{.question}}
We have provided an existing answer: {{.existing_answer}}
We have the opportunity to refine the existing answer
(only if needed) with some more context below.
------------
{{.context}}
------------
INSTRUCTIONS:
- The format you return: Make sure to cite results using [number] notation after the reference. But don't list the references.
Given the new context, refine the original answer to better answer the question. 
If the context isn't useful, return the original answer.`,
		[]string{"question", "existing_answer", "context"},
	)

	qaChain := NewRefineDocuments(
		NewLLMChain(llm, questionPrompt),
		NewLLMChain(llm, refinePrompt),
	)

	results, err := Call(
		context.Background(),
		qaChain,
		map[string]any{
			"input_documents": docs,
			"question":        "APD机器和治疗贵吗?用中文回复我",
		},
		WithTemperature(0.0),
		WithMaxTokens(2048),
	)
	require.NoError(t, err)

	_, ok := results["text"].(string)
	t.Log(results["text"])
	require.True(t, ok, "result does not contain text key")
}

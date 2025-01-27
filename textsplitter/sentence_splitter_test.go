package textsplitter

import (
	"fmt"
	"github.com/aresa7796/sentences/english"
	"testing"
)

func TestNewSentenceSplitter(t *testing.T) {
	splitter := NewSentenceSplitter(150, 20)
	text := `2017ECNP\|为什么女人没有表现的更像男人？
2017-10-10 灵北学院
6
↗如果您喜欢本篇内容，
可以点击右上角菜单分享或转发给好友。
(图片来源于网络)
请注意性别
雅典希腊的克里斯蒂娜达拉在多年基于动物模型的临床前研究中观察到了对抗抑郁药的应答有性别差异并对此提出了自己的见解：具有压力和抑郁特征的实验动物对抗抑郁药应答的性别差异需引起注意。新型和现有抗抑郁药的作用机制需要在雌性和雄性动物的模型中仔细验证。
FST-女性比男性更容易受到压力
例如，在常用的筛选新型抗抑郁药物的测试中，强迫游泳试验(FST)—雌性动物的行为表现与雄性不同。与雄性相比，雌性大鼠表现出更低的活动性。在FST和慢性轻度压力测试之后，人们也注意到大脑中化学物质在不同性别之间的差异。与雄性相比，雌性大鼠表现出海马区5-HT激活功能受损，并且还被观察到其额叶皮质区多巴胺水平的变化。
SSRIs可以减少雌、雄性动物间应激反应的差异，但雌性需要的剂量比雄性低。虽然啮齿动物不是人类，但这确实表明使用这种区分性别的压力测试评估后，女性患者采用SSRIs或任何药物时，可能需要改变剂量。
神经类固醇-神经解剖学的性别差异表达
在大脑中产生的类固醇—神经甾体—起着神经发育和神经可塑性的作用，从而在生理和病理学上产生性别特异性差异。西班牙马德里的LuisCarciaSegura，概述了大脑中神经甾体生成的生物化学阶段，这是一种独立于外周类固醇生成调节的过程。血浆类固醇水平并不代表大脑中类固醇水平，并且在睾丸切除术或卵巢切除术后，大脑适应了性别和区域特异性的神经类固醇水平。
神经类固醇病理学差异对糖尿病的影响
在糖尿病动物模型中可以观察到病理学差异对神经甾体发挥其作用产生影响。性别差异与坐骨神经中神经甾体脱氢表雄酮(DHEA)和雌二醇的水平相关。在人类中，男性经历神经病变的频率更高，而女性经历神经性疼痛和负面感知症状的情况更多。令人感兴趣的是，在动物模型中观察到的神经类固醇效应是否与人类中的不同。
表观遗传学—基因表达显示性别依赖性
采用表观遗传学技术研究了对应激敏感性调节的性别差异。弗吉尼亚州乔治亚•霍德斯解释了DNA甲基化对成年动物性别特异性压力脆弱性的影响。将两种性别的小鼠进行变应力(VS)测试，然后检查其行为以确定其敏感程度。在第6天，雌性小鼠比雄性小鼠更敏感。直到第21天，雄性老鼠才屈服。免疫组织化学研究表明，雌性小鼠具有电路特异性突触前改变，可能促进了其早期的应激敏感性。
他们还注意到对第6天和第21天VS响应的转录激活的性别特异性模式。在伏隔核中操纵DNA甲基转移的甲基转移酶3a(Dnmt3a)的水平揭示了VS实验中的行为差异与性别差异。Dnmt3a表达的增加诱导了两种性别的应激易感性。通过切除Dnmt3a，可使雌性转录组转变为更像雄性的转录模式，同时也提高了其抗压能力。
令人兴奋的是，抑郁症患者中也会出现基因转录激活的性别特异性模式。
目前，Hodes博士正在进行动物研究，通过操纵Dnmt3a，研究雄性转录组是否可以在21天内转变为更像雌性样的转录模式。实际上她正在尝试着建立在变应力中，“男人”是否可以更像“女人”的模式！`
	texts := splitter.SplitText(text)

	for i, t := range texts {
		fmt.Printf("chunk: %d,len: %d, content: %s\n\n", i, len(splitter.TokenEncode(t)), t)
	}
}

func TestSentence(t *testing.T) {
	text := `2017ECNP\|为什么女人没有表现的更像男人？
2017-10-10 灵北学院
6
↗如果您喜欢本篇内容，
可以点击右上角菜单分享或转发给好友。
(图片来源于网络)
请注意性别
雅典希腊的克里斯蒂娜达拉在多年基于动物模型的临床前研究中观察到了对抗抑郁药的应答有性别差异并对此提出了自己的见解：具有压力和抑郁特征的实验动物对抗抑郁药应答的性别差异需引起注意。新型和现有抗抑郁药的作用机制需要在雌性和雄性动物的模型中仔细验证。
FST-女性比男性更容易受到压力
例如，在常用的筛选新型抗抑郁药物的测试中，强迫游泳试验(FST)—雌性动物的行为表现与雄性不同。与雄性相比，雌性大鼠表现出更低的活动性。在FST和慢性轻度压力测试之后，人们也注意到大脑中化学物质在不同性别之间的差异。与雄性相比，雌性大鼠表现出海马区5-HT激活功能受损，并且还被观察到其额叶皮质区多巴胺水平的变化。
SSRIs可以减少雌、雄性动物间应激反应的差异，但雌性需要的剂量比雄性低。虽然啮齿动物不是人类，但这确实表明使用这种区分性别的压力测试评估后，女性患者采用SSRIs或任何药物时，可能需要改变剂量。
神经类固醇-神经解剖学的性别差异表达
在大脑中产生的类固醇—神经甾体—起着神经发育和神经可塑性的作用，从而在生理和病理学上产生性别特异性差异。西班牙马德里的LuisCarciaSegura，概述了大脑中神经甾体生成的生物化学阶段，这是一种独立于外周类固醇生成调节的过程。血浆类固醇水平并不代表大脑中类固醇水平，并且在睾丸切除术或卵巢切除术后，大脑适应了性别和区域特异性的神经类固醇水平。
神经类固醇病理学差异对糖尿病的影响
在糖尿病动物模型中可以观察到病理学差异对神经甾体发挥其作用产生影响。性别差异与坐骨神经中神经甾体脱氢表雄酮(DHEA)和雌二醇的水平相关。在人类中，男性经历神经病变的频率更高，而女性经历神经性疼痛和负面感知症状的情况更多。令人感兴趣的是，在动物模型中观察到的神经类固醇效应是否与人类中的不同。
表观遗传学—基因表达显示性别依赖性
采用表观遗传学技术研究了对应激敏感性调节的性别差异。弗吉尼亚州乔治亚•霍德斯解释了DNA甲基化对成年动物性别特异性压力脆弱性的影响。将两种性别的小鼠进行变应力(VS)测试，然后检查其行为以确定其敏感程度。在第6天，雌性小鼠比雄性小鼠更敏感。直到第21天，雄性老鼠才屈服。免疫组织化学研究表明，雌性小鼠具有电路特异性突触前改变，可能促进了其早期的应激敏感性。
他们还注意到对第6天和第21天VS响应的转录激活的性别特异性模式。在伏隔核中操纵DNA甲基转移的甲基转移酶3a(Dnmt3a)的水平揭示了VS实验中的行为差异与性别差异。Dnmt3a表达的增加诱导了两种性别的应激易感性。通过切除Dnmt3a，可使雌性转录组转变为更像雄性的转录模式，同时也提高了其抗压能力。
令人兴奋的是，抑郁症患者中也会出现基因转录激活的性别特异性模式。
目前，Hodes博士正在进行动物研究，通过操纵Dnmt3a，研究雄性转录组是否可以在21天内转变为更像雌性样的转录模式。实际上她正在尝试着建立在变应力中，“男人”是否可以更像“女人”的模式！`

	tokenizer, err := english.NewSentenceTokenizer(nil)
	if err != nil {
		panic(err)
	}

	sentences := tokenizer.Tokenize(text)
	for i, s := range sentences {
		fmt.Printf("Chunk:[%d],Content:[%s]\n\n ", i, s.Text)
	}
}

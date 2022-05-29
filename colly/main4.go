package main

import (
	"ArSearch/pkg/service"
	"ArSearch/pkg/service/service_schema"

	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

//==========> 130646
/**
{
    "content":{
        "body":"## 盘点国内知名的三位加密艺术家之一--蔡国强\n\n![](https://images.mirror-media.xyz/publication-images/_dzsSqRmi-p6JuufBuwk3.jpg?height=1105\u0026width=2000)\n\n伴随着加密艺术的火热，国内出现了一批NFT艺术领域的“实践者”们，他（她）们具有独特的眼光，大胆的，新潮的思维，愿意去尝试使用NFT，去诞生出不一样的艺术。\n\n区块链技术让数字商品\\*\\*可确权，易流转，可追溯，不可篡改，\\*\\*为发展数字经济打下了坚实的基础。数字艺术是指通过数码设备进行创作或者加工的艺术作品。一方面使用新的媒介和手段创作或复现已有的艺术形式的作品，另一方面使用媒介或者创作手段创作出全新的艺术形式，该形式主要由计算机或算法决定。\n\n从呈现形式以及技术手段分类，数字艺术主要包括以下形式：数字影像、互动装置、虚拟现实、多媒体，卡通动漫，数字游戏，网络艺术，数字设计，动画，数字摄影，数字音乐等。也就是说，数字艺术是艺术和科技高度融合的多学科的交叉领域，涵盖了艺术、科技、文化、教育等诸多方面的内容。一切由数字技术制作的媒体文化，且具有独立的审美价值的都可归属于数字艺术的范畴。\n\n相对于传统艺术作品来说，数字艺术具有**多元性、虚拟性、便捷性、交互性**等特性，具体表现如下：\n\n多元性：数字艺术不再拘泥于传统的绘画形式以及现实世界中的材料的物理限制，可以使用数字手段复现现实或者已有的艺术形式，还可以通过计算机算法生成自动艺术等多种创作形式，是一个集合了自然、人文和社会科学的综合性学科，数字艺术突破时间和空间的限制，使得艺术表达更具有广泛性和多元性。快捷性：数字艺术是由 0 和 1 的字节组成虚拟艺术形式，存在于虚拟的世界中，它没有物理世界中的载体，可以在互联网中进行快速的传播，只要一个上传按键，便可以让全世界共同来欣赏该艺术作品。数字艺术的传递具有快捷性，避免了传统世界中搬运展示等问题。交互性：数字艺术打破传统的艺术展示形式，将原先单一的静态展示的方式改为了可体验、可交互的多重体验方式，观者的欣赏角度也从平行线改为了交叉线，进入到艺术作品中，成为作品创作的一个输入部分，与艺术家一起共同完成创作，感受作品的维度更加广泛且深入。\n\n虚拟性：数字艺术是由字节组成，以数据的形式存储在虚拟的空间内。存储和传播都借助互联网技术，实现在虚拟世界中的快速转换、传播、存储和复制。同时该属性也使得数字艺术需要依赖于计算机和网络等电子设备进行。\n\n加密艺术得到迅速发展，吸引了众多的传统艺术家和收藏家，跳入这个“兔子洞”。国内也有许多艺术家涌入这个领域，下面请让我们一起盘点国内的知名的加密艺术家--**蔡国强**。\n\n![](https://images.mirror-media.xyz/publication-images/ZB6_R7h550ls_M7LuBkgJ.png?height=325\u0026width=403)\n\n我想你一定听过某中国艺术家的烟花表演，他的艺术烟花精彩纷呈，没错他就是--蔡国强。以下是援引来自百度百科的介绍：蔡国强于1957年生于福建泉州，1981-1985年就读于[上海戏剧学院](https://baike.baidu.com/item/%E4%B8%8A%E6%B5%B7%E6%88%8F%E5%89%A7%E5%AD%A6%E9%99%A2)舞台美术系。他的艺术表现横跨绘画、装置、录像及表演艺术等数种媒材。1986-1995年底旅居日本期间，他持续探索从家乡泉州开始的以火药创作绘画的艺术手法，逐渐推大其作品的爆破规模和形式，并建立随后著名的室外爆破计划。他以东方哲学和当代社会问题作为作品观念的根基，因地制宜，阐释和回应当地文化历史。他以艺术的力量和强悍的作品视觉漫步全球，体现在不同文化里自由往来的游牧精神。**他著名的火药爆破艺术和大型装置充满活力和爆发力，超越平面，从室内空间走入社会和自然。**\n\n![](https://images.mirror-media.xyz/publication-images/HArTvh5COb8dopA5w2icJ.png?height=438\u0026width=660)\n\n三十多年来，他的艺术足迹几乎遍及五大洲所有国际大展，在众多世界著名的艺术殿堂举办个展，包括2006年纽约大都会博物馆个展《蔡国强在屋顶：透明纪念碑》和2008年纽约古根海姆美术馆个人回顾展《我想要相信》。他还担任2008年北京奥运会开闭幕式核心创意小组成员及视觉特效艺术总设计。2013 年10月，他在巴黎卢浮宫、奥赛美术馆之间的塞纳河上，创作观念爆破计划《一夜情》。同年，个展《农民达芬奇》在巴西三个城市巡回，吸引观众100 万余人，其中里约热内卢一站成为当年全球所有在世艺术家观展人数最高展览。他的爆破计划《天梯》于2015年6月15日凌晨在福建泉州惠屿岛港湾实现。相关纪录片2016年在美国首映。\n\n![其故乡小渔村实现的《天梯》（2015）](https://images.mirror-media.xyz/publication-images/J-RmjuBdG2t8uzfAPpG82.png?height=493\u0026width=725)\n\n作品收藏机构包括：纽约现代美术馆、大都会博物馆、古根海姆美术馆，巴黎蓬皮杜美术馆，伦敦泰特美术馆、卡塔尔博物馆管理局等世界各国重要博物馆和艺术机构。\n\n蔡国强老师在非加密领域拥有非常多的成就，下面介绍它的加密历程。2021年蔡国强老师的首件NFT作品《瞬间的永恒--101个火药画的引爆》，在TR lab线上平台的48小时拍卖中以250万美元成交。本次由上海外滩美术馆委约创作的《瞬间的永恒》将火药画的重要组成部分——“引爆瞬间”，转化为NFT。其中100个爆破瞬间成为一件NFT，皆来自蔡国强近年“一个人的西方艺术史之旅”的作品创作。\n\n北京时间2021年9月3日上午9点开始在线上平台TR Lab发布自己的第二件NFT作品，名为《炸自己》的99个限量版NFT，销售所得用于支持上海外滩美术馆的教育项目。\n\n![蔡国强，《瞬间的永恒》作品截图，2021年，蔡工作室提供](https://images.mirror-media.xyz/publication-images/sxBukjvjAnURJLyu0S3Ou.png?height=336\u0026width=600)\n\n![蔡国强与《炸自己》诞生的火药画，2021年；罗桑 摄，蔡工作室提供](https://images.mirror-media.xyz/publication-images/XHeOUWm4b_0nmIC6LVh0x.png?height=423\u0026width=679)\n\n蔡国强为《你的白天烟花》NFT项目特别创作手稿，所绘十二种白天烟花，均来自项目中烟花「燃放」的十二个国家，**艺术家将「手稿的创作过程」转化为一件纪念版NFT《想象这样的》**。\n\n一直以来，手稿是蔡老师各种艺术项目探索灵感和贯穿创作过程的重要手段，尤其是他创作烟花爆破计划的重要环节：常常边画手稿，边向烟花厂家和技术人员，以及燃放所在地政府、项目主办方和支持者等各方，说明烟花的技术开发、艺术效果及编排。\n\n![](https://images.mirror-media.xyz/publication-images/sxKEP2tOX4LTDLZbJaZ6s.png?height=508\u0026width=831)\n\n通过《想象这样的》这件特别作品，蔡期待与大家分享更多鲜为公众所知的创作过程，更借此感激社区和藏家一直以来对《你的白天烟花》的支持。\n\n《想象这样的》于2022年5月4日发表，恰逢蔡国强在上海外滩美术馆开馆展「农民达芬奇」十二周年纪念日，也以此作品见证蔡与TRLab联合创始人兼CEO Audrey Ou和美术馆创始人当年结下的友谊。\n\n![](https://images.mirror-media.xyz/publication-images/pOrXdmv3zk8btWGbnmFKC.png?height=443\u0026width=831)\n\n蔡国强老师在阐释《瞬间的永恒——101个火药画的引爆》的创作时也提及了他对NFT这一新媒介的“去中心化”的理解。他说，尽管NFT加密艺术的出现让我们传统意义上所依赖的一些物理空间、物质材料还有交易方法的理解都产生了改变，这是其所谓“去中心化”的一点。但是它也有可能建立另外一种“中心化”，比如它还是一个小圈子。我们其实只是一直在玩一个游戏，虽然某种意义上其形式可能打破了固有艺术形式的中心化，但是也不等于这个物质空间的媒介和教育方法的打破。在蔡国强看来，NFT这一新媒介有着窥视另一个世界的可能，而这种可能充分映照了他对于艺术追求的核心——如何用“看得见”表现“看不见”的世界。因此，这一新作品的突破与尝试也是蔡国强对一些看不见世界的能量释放。但是这仅仅也只是一个开始，之后，他会一步一步地去描述它。",
        "timestamp":1652406160,
        "title":"氪氚科普🐨第四期"
    },
    "digest":"qY9IEsN4zBcoBlfDML7BQToqRNTEbtrn6WVhYv2Dm0Y",
    "authorship":{
        "contributor":"0xc5E2Fd0bC81CA7197fbF3f65bceF1dbc45AaC2Fd",
        "signingKey":"{\"crv\":\"P-256\",\"ext\":true,\"key_ops\":[\"verify\"],\"kty\":\"EC\",\"x\":\"wUr9OmVcfZjjfIJA9NKhYt-KkWXTmxItTnv63JoJ-m8\",\"y\":\"s0XcxmiGGRLHpRTGcG8uzHd4jqkX2VwPVIHJoGIji-g\"}",
        "signature":"AnLkXgIDmvfF7GrcYUrMJzxtLr638BAdDXrB2iCTwK-RWQjwLAvYcSypThLeiYlIfCuhjeNEgwPRJjM1et9WYg",
        "signingKeySignature":"0x81fb5e569ad64429da29c9be394311263af9e79ed95b44988856c3c0e83612dd6c964418e375e42210e85b0e0caef06a63981ea429b65eb057be8ea03246ea2e1b",
        "signingKeyMessage":"I authorize publishing on mirror.xyz from this device using:\n{\"crv\":\"P-256\",\"ext\":true,\"key_ops\":[\"verify\"],\"kty\":\"EC\",\"x\":\"wUr9OmVcfZjjfIJA9NKhYt-KkWXTmxItTnv63JoJ-m8\",\"y\":\"s0XcxmiGGRLHpRTGcG8uzHd4jqkX2VwPVIHJoGIji-g\"}",
        "algorithm":{
            "name":"ECDSA",
            "hash":"SHA-256"
        }
    },
    "nft":{

    },
    "version":"04-19-2022",
    "originalDigest":"8iUB7LqSCSjojdXhXsw_8jeMh__77SAZ-l7RODqlLLY",
    "arweave_tx":"nGmXa-K4NJGgu7l4bVP-P9pBqnCo_wpKKBKmq4hUN6E"
}
*/
func main() {
	kafkaURL := "localhost:9092"
	//kafkaURL := "45.76.151.181:9092"
	topic := "arsearch-topic"
	groupID := "groupId"

	ctx := context.Background()
	reader := service.GetKafkaReader(kafkaURL, topic, groupID)
	wg := sync.WaitGroup{}

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("recover err==>", err)
		}
	}()

	group, _ := errgroup.WithContext(ctx)
	for {
		wg.Add(1)

		group.Go(func() error {
			m, err := reader.ReadMessage(context.Background())
			if err != nil {
				log.Println(err)
			}

			v := service_schema.ArData{}
			json.Unmarshal(m.Value, &v)

			link1 := fmt.Sprintf("https://mirror.xyz/%s/%s", v.Authorship.Contributor, v.OriginalDigest)
			mirroData := service_schema.MirrorData{
				Id:             m.Offset,
				Title:          v.Content.Title,
				Content:        v.Content.Body,
				Digest:         v.Digest,
				Link:           link1,
				OriginalDigest: v.OriginalDigest,
				CreatedAt:      time.Unix(int64(v.Content.Timestamp), 0),
				PublishedAt:    time.Unix(int64(v.Content.Timestamp), 0),
				ArweaveTx:      v.ArWeaveTx,
				Contributor:    v.Authorship.Contributor,
			}
			data, err1 := service.SaveMirrorData(&mirroData)
			if err1 != nil {
				fmt.Println("err===>", err1)
			}
			fmt.Println("data=====>", data)
			wg.Done()
			return err1
		})

		//time.Sleep(time.Millisecond * 200)
		//fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		//time.Sleep(time.Second * 50)
	}
	wg.Wait()
}

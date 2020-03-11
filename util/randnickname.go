package util

import "math/rand"

var NickNames = []string{"月色如浅梦", "临风纵饮", "凡尘红梦", "子兮分袂", "墨雨云烟", "白首迟暮", "二十四桥明月夜", "怀揽清风", "故里归长安", "酒肆饮几壶", "晨风韵雨", "晚烟许清扇", "雪蝶宿秋风", "梅坼筱枫", "枯蝶残香零叶", "商城凤笙", "回首闻君叹", "清冷孤傲古风网名女生", "青衫烟雨客", "似是故人来", "常闻君言笑", "素手挽清风", "半绾青丝注斜阳", "箫曲清歌", "曾经沧海难为水", "胭脂泪几时垂", "倦了轻狂少年", "古巷烟雨゛断桥殇", "花落君独醉", "西楼听雨", "墨雨轩", "月色映归客", "怎知春色几许", "寻欢人", "季末花忆残", "岁莫染三生", "宛音了之", "枫尘于往逝", "栀里墨思量", "青袖沾南风", "心陌南尘", "羽月风花", "叶隐知心魂", "浮世清欢", "一念執著", "落日映苍穹", "流苏复流苏", "江畔秋时月", "山河不入梦", "青黛", "明月清风", "云深不知处", "北城柳絮飘", "墨染傾城ゞ", "雁字回时，月满西楼"}

var Avatar = []string{"avatar0.png",
	"avatar1.jpg",
	"avatar2.jpg",
	"avatar3.jpg",
	"avatar4.jpg",
	"avatar5.jpg"}

func RandNickname() string {
	randIndex := rand.Int31n(int32(len(NickNames)))
	return NickNames[randIndex]
}

func RandAvatar() string {
	randIndex := rand.Int31n(int32(len(Avatar)))
	return Avatar[randIndex]
}

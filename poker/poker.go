package grpc_demo

import "errors"

// PokerTransferList
// rank, rank->prime, rank->bit, suit
var PokerTransferList = map[string][4]int{
	"2d": {0, 2, 65536, 16384},         // Diamond 2
	"3d": {256, 3, 131072, 16384},      // Diamond 3
	"4d": {512, 5, 262144, 16384},      // Diamond 4
	"5d": {768, 7, 524288, 16384},      // Diamond 5
	"6d": {1024, 11, 1048576, 16384},   // Diamond 6
	"7d": {1280, 13, 2097152, 16384},   // Diamond 7
	"8d": {1536, 17, 4194304, 16384},   // Diamond 8
	"9d": {1792, 19, 8388608, 16384},   // Diamond 9
	"Td": {2048, 23, 16777216, 16384},  // Diamond 10
	"Jd": {2304, 29, 33554432, 16384},  // Diamond J
	"Qd": {2560, 31, 67108864, 16384},  // Diamond Q
	"Kd": {2816, 37, 134217728, 16384}, // Diamond K
	"Ad": {3072, 41, 268435456, 16384}, // Diamond A
	"2c": {0, 2, 65536, 32768},         // Club 2
	"3c": {256, 3, 131072, 32768},      // Club 3
	"4c": {512, 5, 262144, 32768},      // Club 4
	"5c": {768, 7, 524288, 32768},      // Club 5
	"6c": {1024, 11, 1048576, 32768},   // Club 6
	"7c": {1280, 13, 2097152, 32768},   // Club 7
	"8c": {1536, 17, 4194304, 32768},   // Club 8
	"9c": {1792, 19, 8388608, 32768},   // Club 9
	"Tc": {2048, 23, 16777216, 32768},  // Club 10
	"Jc": {2304, 29, 33554432, 32768},  // Club J
	"Qc": {2560, 31, 67108864, 32768},  // Club Q
	"Kc": {2816, 37, 134217728, 32768}, // Club K
	"Ac": {3072, 41, 268435456, 32768}, // Club A
	"2h": {0, 2, 65536, 8192},          // Heart 2
	"3h": {256, 3, 131072, 8192},       // Heart 3
	"4h": {512, 5, 262144, 8192},       // Heart 4
	"5h": {768, 7, 524288, 8192},       // Heart 5
	"6h": {1024, 11, 1048576, 8192},    // Heart 6
	"7h": {1280, 13, 2097152, 8192},    // Heart 7
	"8h": {1536, 17, 4194304, 8192},    // Heart 8
	"9h": {1792, 19, 8388608, 8192},    // Heart 9
	"Th": {2048, 23, 16777216, 8192},   // Heart 10
	"Jh": {2304, 29, 33554432, 8192},   // Heart J
	"Qh": {2560, 31, 67108864, 8192},   // Heart Q
	"Kh": {2816, 37, 134217728, 8192},  // Heart K
	"Ah": {3072, 41, 268435456, 8192},  // Heart A
	"2s": {0, 2, 65536, 4096},          // Spade 2
	"3s": {256, 3, 131072, 4096},       // Spade 3
	"4s": {512, 5, 262144, 4096},       // Spade 4
	"5s": {768, 7, 524288, 4096},       // Spade 5
	"6s": {1024, 11, 1048576, 4096},    // Spade 6
	"7s": {1280, 13, 2097152, 4096},    // Spade 7
	"8s": {1536, 17, 4194304, 4096},    // Spade 8
	"9s": {1792, 19, 8388608, 4096},    // Spade 9
	"Ts": {2048, 23, 16777216, 4096},   // Spade 10
	"Js": {2304, 29, 33554432, 4096},   // Spade J
	"Qs": {2560, 31, 67108864, 4096},   // Spade Q
	"Ks": {2816, 37, 134217728, 4096},  // Spade K
	"As": {3072, 41, 268435456, 4096},  // Spade A
}

var handsCombination = [][]int{
	{0, 1, 2, 3, 4}, {0, 1, 2, 3, 5}, {0, 1, 2, 3, 6},
	{0, 1, 2, 4, 5}, {0, 1, 2, 4, 6}, {0, 1, 2, 5, 6},
	{0, 1, 3, 4, 5}, {0, 1, 3, 4, 6}, {0, 1, 3, 5, 6},
	{0, 1, 4, 5, 6}, {0, 2, 3, 4, 5}, {0, 2, 3, 4, 6},
	{0, 2, 3, 5, 6}, {0, 2, 4, 5, 6}, {0, 3, 4, 5, 6},
	{1, 2, 3, 4, 5}, {1, 2, 3, 4, 6}, {1, 2, 3, 5, 6},
	{1, 2, 4, 5, 6}, {1, 3, 4, 5, 6}, {2, 3, 4, 5, 6},
}

func PokerEvaluator(self, board []string) (string, error) {
	self = append(self, board...)
	var cards [][4]int

	// 七張牌
	for i := 0; i < len(self); i++ {
		cards = append(cards, PokerTransferList[self[i]])
	}

	// 組成五張牌一組的所有類型的手牌，C7取5種組合
	allHands := make([][][4]int, len(handsCombination))
	for k1, i1 := range handsCombination {
		allHands[k1] = make([][4]int, 5)
		for k2, i2 := range i1 {
			allHands[k1][k2] = cards[i2]
		}
	}

	bestScore := 7462
	for _, hand := range allHands {
		product := 1
		var hands []int
		for _, item := range hand {
			hands = append(hands, item[0]|item[1]|item[2]|item[3])
			// 五張牌質數相乘
			product *= item[1]
		}
		score := getScore(hands, product)
		if score < bestScore {
			bestScore = score
		}
	}
	return scoreToType(bestScore)
}

func getScore(hands []int, product int) int {
	if hands[0]&hands[1]&hands[2]&hands[3]&hands[4]&0xF000 > 0 {
		// 同花
		return flushLookup[product]
	}
	// 非同花
	return nonFlushLookup[product]
}

func scoreToType(score int) (s string, err error) {
	switch {
	case score == 1:
		s = "Royal Flush" // 皇家同花順
	case score <= 10 && score >= 2:
		s = "Straight Flush" // 同花順
	case score <= 166 && score >= 11:
		s = "Four of a Kind" // 四條
	case score <= 322 && score >= 167:
		s = "Full house" // 葫蘆
	case score <= 1599 && score >= 323:
		s = "Flush" // 同花
	case score <= 1609 && score >= 1600:
		s = "Straight" // 順子
	case score <= 2467 && score >= 1610:
		s = "Three of a kind" // 三條
	case score <= 3325 && score >= 2468:
		s = "Two Pairs" // 兩對
	case score <= 6185 && score >= 3326:
		s = "One Pair" // 一對
	case score <= 7462 && score >= 6186:
		s = "High card" // 高牌
	default:
		s = "WTF.." // 看到鬼
		err = errors.New("input error")
	}
	return
}

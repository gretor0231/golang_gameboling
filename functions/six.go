package functions

//calculate win rate at river
//Complete the function f that computes the ranking of a poker hand. The input is consisted of two slices, the first slice contains two integers, the second slice contains five integers. Each integer represents different poker cards.
//Belows are the correspoding relationship between each interger(for better reading, the integer is shown in hexadecimal) and poker card:
//
//0x102,0x103,0x104,0x105,0x106,0x107,0x108,0x109,0x10a,0x10b,0x10c,0x10d,0x10e correspond to 2,3,4,5,6,7,8,9,10,J,Q,K,A of diamond.
//0x202,0x203,0x204,0x205,0x206,0x207,0x208,0x209,0x20a,0x20b,0x20c,0x20d,0x20e correspond to 2,3,4,5,6,7,8,9,10,J,Q,K,A of club.
//0x302,0x303,0x304,0x305,0x306,0x307,0x308,0x309,0x30a,0x30b,0x30c,0x30d,0x30e correspond to 2,3,4,5,6,7,8,9,10,J,Q,K,A of heart.
//0x402,0x403,0x404,0x405,0x406,0x407,0x408,0x409,0x40a,0x40b,0x40c,0x40d,0x40e correspond to 2,3,4,5,6,7,8,9,10,J,Q,K,A of spade.
//
//Find the min rank from all five cards combination.
//You only have one opponent. Calcualte the percentage of you can win, cut the decimal parts.
//
//Please offer a verison with the better performance.

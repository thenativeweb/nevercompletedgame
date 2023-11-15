package game

type level struct {
	Question string
	Answer   string
	//Hint     string
}

var levels = []level{
	{
		// 01
		Question: "First letter of the alphabet.",
		Answer:   "A",
		//Hint:     "This is not a trick - question",
	}, {
		// 02
		Question: "2 + 1 = ?",
		Answer:   "3",
		//Hint:     "Simple math",
	}, {
		// 03
		Question: "A, B, C, D, E, _?",
		Answer:   "F",
		//Hint:     "Fill in the blank",
	}, {
		// 04
		Question: "5 x 7 = ? ",
		Answer:   "35",
		//Hint:     "More math",
	},
	{
		// 05
		Question: "We'll be counting stars ********",
		Answer:   "8",
		//Hint:     "Count the stars",
	},
	{
		// 06
		Question: "f4nt45t1c",
		Answer:   "fantastic",
		//Hint:     "Leetspeak",
	},
	{
		// 07
		Question: "A5 B4 C3 D2 __.",
		Answer:   "E1",
		//Hint:     "A series brain teaser",
	},
	{
		// 08
		Question: "Roman XV",
		Answer:   "15",
		//Hint:     "You know your roman numerals, right?",
	},

	{
		// 09
		Question: "Lightens the night.",
		Answer:   "moon",
		//Hint:     "Look up",
	},
	{
		// 10
		Question: "Forever and ever. ∞",
		Answer:   "infinity",
		//Hint:     "What is the symbol?",
	},
	{
		// 11
		Question: "British Color",
		Answer:   "colour",
		//Hint:     "They spell differently",
	},
	{
		// 12
		Question: "2, 3, 5, 7, 11, 13, 17 What are they called?",
		Answer:   "prime numbers",
		//Hint:     "The next one is 19",
	},
	{
		// 13
		Question: "3.14159265...",
		Answer:   "pi",
		//Hint:     "Is a greek letter",
	},

	{
		// 14
		Question: "1 + 2 + 3 + ... + 999 + 1000 ",
		Answer:   "500500",
		//Hint:     "Simple Gauss summation",
	},
	{
		// 15
		Question: "123567890134567890123456790",
		Answer:   "428",
		//Hint:     "What is missing",
	},
	{
		// 16
		Question: "The answer is <span style=\"opacity: 0.0;\">hidden</span>.",
		Answer:   "hidden",
		//Hint:     "Copy and Paste",
	},
	{
		// 17
		Question: "1729 = 1^3 + 12^3 = 9^3 + 10^3",
		Answer:   "Ramanujan",
		//Hint:     "Google 1729",
	},
	{
		// 18
		Question: "- .-. .- -. ... .-.. .- - .. --- -.",
		Answer:   "translation",
		//Hint:     "Morse code",
	},
	{
		// 19
		Question: "---- . ------ ------ --------- --",
		Answer:   "Feigenbaum",
		//Hint:     "Count and google",
	},
	{
		// 20
		Question: "Fired a round, now he's deceased.",
		Answer:   "Queen",
		//Hint:     "A modified lyric from a popular band's song",
	},
	{
		// 21
		Question: "//000 T '29",
		Answer:   "Wall Street",
		//Hint:     "A color and a date",
	},

	{
		// 22
		Question: "97...98...99...00",
		Answer:   "millennium bug",
		//Hint:     "Think about dates. Particularly what might have happened on the last one",
	},
	{
		// 23
		Question: "RWluc3RlaW4ncy4",
		Answer:   "relativity",
		//Hint:     "Base64",
	},
	{
		// 24
		Question: "230809200519 152120080518 140709011420",
		Answer:   "antarctica",
		//Hint:     "02=B",
	},
	{
		// 25
		Question: "Stop! Stop! Stop! Start. Stop!\nStop! Start. Stop! Stop! Start.\nStop! Stop! Start. Stop! Stop!\nStop! Stop! Stop! Stop! Stop!\nOK!",
		Answer:   "weleetka",
		//Hint:     "Binary address",
	},
	{
		// 26
		Question: "VlZSS1IyUldiRmhUYldocVltdHdNbE5WV2s5aFIwcDBWVzVhYTJSNk1Eaz0",
		Answer:   "1867",
		//Hint:     "Base64 relation",
	},
	{
		// 27
		Question: "Strong Chase Stone Hunt.",
		Answer:   "justice",
		//Hint:     "The common ground",
	},
	//{
	//// 28
	//Question: [unknown image],
	//Answer:   "Deadly Sins',
	////Hint:     "Zoom in. Caesar",
	//},
	{
		// 29
		Question: "10111100 10010000 10011000 10010110 10001011 10010000 11011111 10011010 10001101 10011000 10010000 11011111 10001100 10001010 100100101",
		Answer:   "descartes",
		//Hint:     "NOT",
	},
	{
		// 30
		Question: "16th A 22nd 35th O 38th E 5th 14th A 45th 23rd Y",
		Answer:   "psychopath",
		//Hint:     "US Presidents",
	},
	//{
	//// 32
	//Question: [unknown video],
	//Answer:   "ted bundy',
	////Hint:     "There are white dots that appear at 2:22 and 4:08.",
	//},
	{
		// 32
		Question: "149 14 97 12 124 3 143 3 194 2 75 5 161 10",
		Answer:   "world trade center",
		//Hint:     "149 mod 14 = 9",
	},
	{
		// 33
		Question: "D73 S04 P86 B07 S95",
		Answer:   "european union",
		//Hint:     "Letters represent countries",
	},
	{
		// 34
		Question: "Try a thing You haven’t done three times. Once, to get over the Fear of doing It. Twice, to learn how to do it. And a third time, to figure out whether You Like It or not.",
		Answer:   "horror movie",
		//Hint:     "Wrongly capitalized words",
	},
	{
		// 34
		Question: "Edxwer Rgnvgy Resxc Iuhnm",
		Answer:   "corpus christi",
		//Hint:     "Shapes on keyboard (QWERTY), City",
	},
	{
		// 34
		Question: "L2:7\n1,799,160,000,000 km\n1,822,660,000,000 km",
		Answer:   "world war",
		//Hint:     "Ratios",
	},
	{
		// 34
		Question: "Wourali mal cam sis.\n1331471115165824910651",
		Answer:   "plato",
		//Hint:     "Numbers index into sentence",
	},
	{
		// 34
		Question: "Queal dan pope itos creation.",
		Answer:   "isaac newton",
		//Hint:     "anagram",
	},
	{
		// 34
		Question: "H4sIAAAAAAAA//Px8XFxcQkKCgKSPj4+XC5gwAUXQMhAWBBOKAiAtAEAy5kSxD8AAAA=",
		Answer:   "johann sebastian",
		//Hint:     "Base64 gzip",
	},
	{
		// 34
		Question: "1580.63 1254.55 1054.94 627.27 939.85 837.31 41.902 12.454",
		Answer:   "borghese",
		//Hint:     "musical frequencies, coordinates, inscription",
	},
	//{
	//// 41
	//Question: "he seventeenth of the primes.\nIts twin is twentynine.\nPo.\nLet’s drive.\nBye, Alfonsín.\nAdd seven triangular.\nSweet.\nIt named the Valley.",
	//Answer:   "???',
	////Hint:     "???",
	//},
}

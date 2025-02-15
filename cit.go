package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	//"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type QA struct {
	Question string
	Answers  []string
}

var questions = []QA{
	{Question: "What is the supreme law of the land?", Answers: []string{
		"the Constitution",
	}},
	{Question: "What does the Constitution do?", Answers: []string{
		"sets up the government",
		"defines the government",
		"protects basic rights of Americans",
	}},
	{Question: "The idea of self-government is in the first three words of the Constitution.What are these words?", Answers: []string{
		"We the People",
	}},
	{Question: "What is an amendment?", Answers: []string{
		"a change (to the Constitution)",
		"an addition (to the Constitution)",
	}},
	{Question: "What do we call the first ten amendments to the Constitution?", Answers: []string{
		"the Bill of Rights",
	}},
	{Question: "What is one right or freedom from the First Amendment?*", Answers: []string{
		"speech",
		"religion",
		"assembly",
		"press",
		"petition the government",
	}},
	{Question: "How many amendments does the Constitution have?", Answers: []string{
		"twenty-seven (27)",
	}},
	{Question: "What did the Declaration of Independence do?", Answers: []string{
		"announced our independence (from Great Britain)",
		"declared our independence (from Great Britain)",
		"said that the United States is free (from Great Britain)",
	}},
	{Question: "What are two rights in the Declaration of Independence?", Answers: []string{
		"life",
		"liberty",
		"pursuit of happiness",
	}},
	{Question: "What is freedom of religion?", Answers: []string{
		"You can practice any religion, or not practice a religion.",
	}},
	{Question: "What is the economic system in the United States?*", Answers: []string{
		"capitalist economy",
		"market economy",
	}},
	{Question: "What is the “rule of law”?", Answers: []string{
		"Everyone must follow the law.",
		"Leaders must obey the law.",
		"Government must obey the law.",
		"No one is above the law.",
	}},
	{Question: "Name one branch or part of the government.*", Answers: []string{
		"Congress / legislative",
		"President / executive",
		"the courts / judicial",
	}},
	{Question: "What stops one branch of government from becoming too powerful?", Answers: []string{
		"checks and balances",
		"separation of powers",
	}},
	{Question: "Who is in charge of the executive branch?", Answers: []string{
		"the President",
	}},
	{Question: "Who makes federal laws?", Answers: []string{
		"Congress",
		"Senate and House (of Representatives)",
		"(U.S. or national) legislature",
	}},
	{Question: "What are the two parts of the U.S. Congress?*", Answers: []string{
		"the Senate and House (of Representatives)",
	}},
	{Question: "How many U.S. Senators are there?", Answers: []string{
		"one hundred (100)",
	}},
	{Question: "We elect a U.S. Senator for how many years?", Answers: []string{
		"six (6)",
	}},
	{Question: "Who is one of your state’s U.S. Senators now?", Answers: []string{
		"Adam B. Schiff",
		"Alex Padilla",
	}},
	{Question: "The House of Representatives has how many voting members?", Answers: []string{
		"four hundred thirty-five (435)",
	}},
	{Question: "We elect a U.S. Representative for how many years?", Answers: []string{
		"two (2)",
	}},
	{Question: "Name your U.S. Representative.", Answers: []string{
		"Ro Khanna",
	}},
	{Question: "Who does a U.S. Senator represent?", Answers: []string{
		"all people of the state",
	}},
	{Question: "Why do some states have more Representatives than other states?", Answers: []string{
		"(because of) the state’s population",
		"(because) they have more people",
		"(because) some states have more people",
	}},
	{Question: "We elect a President for how many years?", Answers: []string{
		"four (4)",
	}},
	{Question: "In what month do we vote for President?*", Answers: []string{
		"November",
	}},
	{Question: "What is the name of the President of the United States now?*", Answers: []string{
		"Donald J. Trump",
	}},
	{Question: "What is the name of the Vice President of the United States now?", Answers: []string{
		"JD Vance",
	}},
	{Question: "If the President can no longer serve, who becomes President?", Answers: []string{
		"the Vice President",
	}},
	{Question: "If both the President and the Vice President can no longer serve, who becomes President?", Answers: []string{
		"the Speaker of the House",
	}},
	{Question: "Who is the Commander in Chief of the military?", Answers: []string{
		"the President",
	}},
	{Question: "Who signs bills to become laws?", Answers: []string{
		"the President",
	}},
	{Question: "Who vetoes bills?", Answers: []string{
		"the President",
	}},
	{Question: "What does the President’s Cabinet do?", Answers: []string{
		"advises the President",
	}},
	{Question: "What are two Cabinet-level positions?", Answers: []string{
		"Attorney General",
		"Vice President",
		"Secretary of Agriculture",
		"Secretary of Commerce",
		"Secretary of Defense",
		"Secretary of Education",
		"Secretary of Energy",
		"Secretary of Health and Human Services",
		"Secretary of Homeland Security",
		"Secretary of Housing and Urban Development",
		"Secretary of the Interior",
		"Secretary of Labor",
		"Secretary of State",
		"Secretary of Transportation",
		"Secretary of the Treasury",
		"Secretary of Veterans Affairs",
	}},
	{Question: "What does the judicial branch do?", Answers: []string{
		"reviews laws",
		"explains laws",
		"resolves disputes (disagreements)",
		"decides if a law goes against the Constitution",
	}},
	{Question: "What is the highest court in the United States?", Answers: []string{
		"the Supreme Court",
	}},
	{Question: "How many justices are on the Supreme Court?", Answers: []string{
		"nine (9)",
	}},
	{Question: "Who is the Chief Justice of the United States now?", Answers: []string{
		"John G. Roberts, Jr.",
	}},
	{Question: "Under our Constitution, some powers belong to the federal government. What is one power of the federal government?", Answers: []string{
		"to print money",
		"to declare war",
		"to create an army",
		"to make treaties",
	}},
	{Question: "Under our Constitution, some powers belong to the states. What is one power of the states?", Answers: []string{
		"provide schooling and education",
		"provide protection (police)",
		"provide safety (fire departments)",
		"give a driver’s license",
		"approve zoning and land use",
	}},
	{Question: "Who is the Governor of your state now?", Answers: []string{
		"Gavin Newsom",
	}},
	{Question: "What is the capital of your state?*", Answers: []string{
		"Sacramento",
	}},
	{Question: "What are the two major political parties in the United States?*", Answers: []string{
		"Democratic and Republican",
	}},
	{Question: "What is the political party of the President now?", Answers: []string{
		"Republican",
	}},
	{Question: "What is the name of the Speaker of the House of Representatives now?", Answers: []string{
		"James Michael/Mike Johnson",
	}},
	{Question: "There are four amendments to the Constitution about who can vote. Describe one of them.", Answers: []string{
		"Citizens eighteen (18) and older (can vote).",
		"You don’t have to pay (a poll tax) to vote.",
		"Any citizen can vote. (Women and men can vote.)",
		"A male citizen of any race (can vote).",
	}},
	{Question: "What is one responsibility that is only for United States citizens?*", Answers: []string{
		"serve on a jury",
		"vote in a federal election",
	}},
	{Question: "Name one right only for United States citizens.", Answers: []string{
		"vote in a federal election",
		"run for federal office",
	}},
	{Question: "What are two rights of everyone living in the United States?", Answers: []string{
		"freedom of expression",
		"freedom of speech",
		"freedom of assembly",
		"freedom to petition the government",
		"freedom of religion",
		"the right to bear arms",
	}},
	{Question: "What do we show loyalty to when we say the Pledge of Allegiance?", Answers: []string{
		"the United States",
		"the flag",
	}},
	{Question: "What is one promise you make when you become a United States citizen?", Answers: []string{
		"give up loyalty to other countries",
		"defend the Constitution and laws of the United States",
		"obey the laws of the United States",
		"serve in the U.S. military (if needed)",
		"serve (do important work for) the nation (if needed)",
		"be loyal to the United States",
	}},
	{Question: "How old do citizens have to be to vote for President?*", Answers: []string{
		"eighteen (18) and older",
	}},
	{Question: "What are two ways that Americans can participate in their democracy?", Answers: []string{
		"vote",
		"join a political party",
		"help with a campaign",
		"join a civic group",
		"join a community group",
		"give an elected official your opinion on an issue",
		"call Senators and Representatives",
		"publicly support or oppose an issue or policy",
		"run for office",
		"write to a newspaper",
	}},
	{Question: "When is the last day you can send in federal income tax forms?*", Answers: []string{
		"April 15",
	}},
	{Question: "When must all men register for the Selective Service?", Answers: []string{
		"at age eighteen (18)",
		"between eighteen (18) and twenty-six (26)",
	}},
	{Question: "What is one reason colonists came to America?", Answers: []string{
		"freedom",
		"political liberty",
		"religious freedom",
		"economic opportunity",
		"practice their religion",
		"escape persecution",
	}},
	{Question: "Who lived in America before the Europeans arrived?", Answers: []string{
		"American Indians",
		"Native Americans",
	}},
	{Question: "What group of people was taken to America and sold as slaves?", Answers: []string{
		"Africans",
		"people from Africa",
	}},
	{Question: "Why did the colonists fight the British?", Answers: []string{
		"because of high taxes (taxation without representation)",
		"because the British army stayed in their houses (boarding, quartering)",
		"because they didn’t have self-government",
	}},
	{Question: "Who wrote the Declaration of Independence?", Answers: []string{
		"(Thomas) Jefferson",
	}},
	{Question: "When was the Declaration of Independence adopted?", Answers: []string{
		"July 4, 1776",
	}},
	{Question: "There were 13 original states. Name three.", Answers: []string{
		"New Hampshire",
		"Massachusetts",
		"Rhode Island",
		"Connecticut",
		"New York",
		"New Jersey",
		"Pennsylvania",
		"Delaware",
		"Maryland",
		"Virginia",
		"North Carolina",
		"South Carolina",
		"Georgia",
	}},
	{Question: "What happened at the Constitutional Convention?", Answers: []string{
		"The Constitution was written.",
		"The Founding Fathers wrote the Constitution.",
	}},
	{Question: "When was the Constitution written?", Answers: []string{
		"1787",
	}},
	{Question: "The Federalist Papers supported the passage of the U.S. Constitution. Name one of the writers.", Answers: []string{
		"(James) Madison",
		"(Alexander) Hamilton",
		"(John) Jay",
		"Publius",
	}},
	{Question: "What is one thing Benjamin Franklin is famous for?", Answers: []string{
		"U.S. diplomat",
		"oldest member of the Constitutional Convention",
		"first Postmaster General of the United States",
		"writer of “Poor Richard’s Almanac”",
		"started the first free libraries",
	}},
	{Question: "Who is the “Father of Our Country”?", Answers: []string{
		"(George) Washington",
	}},
	{Question: "Who was the first President?*", Answers: []string{
		"(George) Washington",
	}},
	{Question: "What territory did the United States buy from France in 1803?", Answers: []string{
		"the Louisiana Territory",
		"Louisiana",
	}},
	{Question: "Name one war fought by the United States in the 1800s.", Answers: []string{
		"War of 1812",
		"Mexican-American War",
		"Civil War",
		"Spanish-American War",
	}},
	{Question: "Name the U.S. war between the North and the South.", Answers: []string{
		"the Civil War",
		"the War between the States",
	}},
	{Question: "Name one problem that led to the Civil War.", Answers: []string{
		"slavery",
		"economic reasons",
		"states’ rights",
	}},
	{Question: "What was one important thing that Abraham Lincoln did?*", Answers: []string{
		"freed the slaves (Emancipation Proclamation)",
		"saved (or preserved) the Union",
		"led the United States during the Civil War",
	}},
	{Question: "What did the Emancipation Proclamation do?", Answers: []string{
		"freed the slaves",
		"freed slaves in the Confederacy",
		"freed slaves in the Confederate states",
		"freed slaves in most Southern states",
	}},
	{Question: "What did Susan B. Anthony do?", Answers: []string{
		"fought for women’s rights",
		"fought for civil rights",
	}},
	{Question: "Name one war fought by the United States in the 1900s.*", Answers: []string{
		"World War I",
		"World War II",
		"Korean War",
		"Vietnam War",
		"(Persian) Gulf War",
	}},
	{Question: "Who was President during World War I?", Answers: []string{
		"(Woodrow) Wilson",
	}},
	{Question: "Who was President during the Great Depression and World War II?", Answers: []string{
		"(Franklin) Roosevelt",
	}},
	{Question: "Who did the United States fight in World War II?", Answers: []string{
		"Japan, Germany, and Italy",
	}},
	{Question: "Before he was President, Eisenhower was a general. What war was he in?", Answers: []string{
		"World War II",
	}},
	{Question: "During the Cold War, what was the main concern of the United States?", Answers: []string{
		"Communism",
	}},
	{Question: "What movement tried to end racial discrimination?", Answers: []string{
		"civil rights (movement)",
	}},
	{Question: "What did Martin Luther King, Jr. do?*", Answers: []string{
		"fought for civil rights",
		"worked for equality for all Americans",
	}},
	{Question: "What major event happened on September 11, 2001, in the United States?", Answers: []string{
		"Terrorists attacked the United States.",
	}},
	{Question: "Name one American Indian tribe in the United States.", Answers: []string{
		"Cherokee",
		"Navajo",
		"Sioux",
		"Chippewa",
		"Choctaw",
		"Pueblo",
		"Apache",
		"Iroquois",
		"Creek",
		"Blackfeet",
		"Seminole",
		"Cheyenne",
		"Arawak",
		"Shawnee",
		"Mohegan",
		"Huron",
		"Oneida",
		"Lakota",
		"Crow",
		"Teton",
		"Hopi",
		"Inuit",
	}},
	{Question: "Name one of the two longest rivers in the United States.", Answers: []string{
		"Missouri (River)",
		"Mississippi (River)",
	}},
	{Question: "What ocean is on the West Coast of the United States?", Answers: []string{
		"Pacific (Ocean)",
	}},
	{Question: "What ocean is on the East Coast of the United States?", Answers: []string{
		"Atlantic (Ocean)",
	}},
	{Question: "Name one U.S. territory.", Answers: []string{
		"Puerto Rico",
		"U.S. Virgin Islands",
		"American Samoa",
		"Northern Mariana Islands",
		"Guam",
	}},
	{Question: "Name one state that borders Canada.", Answers: []string{
		"Maine",
		"New Hampshire",
		"Vermont",
		"New York",
		"Pennsylvania",
		"Ohio",
		"Michigan",
		"Minnesota",
		"North Dakota",
		"Montana",
		"Idaho",
		"Washington",
		"Alaska",
	}},
	{Question: "Name one state that borders Mexico.", Answers: []string{
		"California",
		"Arizona",
		"New Mexico",
		"Texas",
	}},
	{Question: "What is the capital of the United States?*", Answers: []string{
		"Washington, D.C.",
	}},
	{Question: "Where is the Statue of Liberty?*", Answers: []string{
		"New York (Harbor)",
		"Liberty Island",
		"New Jersey",
		"near New York City",
		"on the Hudson (River).]",
	}},
	{Question: "Why does the flag have 13 stripes?", Answers: []string{
		"because there were 13 original colonies",
		"because the stripes represent the original colonies",
	}},
	{Question: "Why does the flag have 50 stars?*", Answers: []string{
		"because there is one star for each state",
		"because each star represents a state",
		"because there are 50 states",
	}},
	{Question: "What is the name of the national anthem?", Answers: []string{
		"The Star-Spangled Banner",
	}},
	{Question: "When do we celebrate Independence Day?*", Answers: []string{
		"July 4",
	}},
	{Question: "Name two national U.S. holidays.", Answers: []string{
		"New Year’s Day",
		"Martin Luther King, Jr. Day",
		"Presidents’ Day",
		"Memorial Day",
		"Independence Day",
		"Labor Day",
		"Columbus Day",
		"Veterans Day",
		"Thanksgiving",
		"Christmas",
	}},
}

func main() {
	list := flag.Bool("list", false, "List all questions")
	star := flag.Bool("star", false, "Show only starred questions (65 and older)")
	flag.Parse()

	if *list {
		for i, q := range questions {
			fmt.Println(i+1, q.Question)
			for _, a := range q.Answers {
				fmt.Println(" ", a)
			}
		}
		return
	}

	a := app.New()
	w := a.NewWindow("Citizenship Test")

	question := binding.NewString()
	answers := binding.NewStringList()

	wquestion := widget.NewLabelWithData(question)
	wquestion.TextStyle = fyne.TextStyle{Bold: true}

	wanswers := widget.NewListWithData(
		answers,
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(item binding.DataItem, obj fyne.CanvasObject) {
			obj.(*widget.Label).Bind(item.(binding.String))
		},
	)

	if *star {
		var qs []QA
		for _, q := range questions {
			if strings.HasSuffix(q.Question, "*") {
				qs = append(qs, q)
			}
		}
		questions = qs
	}

	qs := rand.Perm(len(questions)) // [0:10]
	current := 0
	total := len(qs)

	title := "Answer 6 questions correctly out of 10 to pass the test    (%v/%v)"
	wtitle := widget.NewLabel(fmt.Sprintf(title, 0, len(qs)))

	w.SetContent(container.NewVBox(
		wtitle,
		wquestion,
		wanswers,
		widget.NewButton("Next Question", func() {
			if current < len(qs) {
				q := questions[qs[current]]
				question.Set(q.Question)
				answers.Set(append([]string{""}, q.Answers...))
				wanswers.ScrollToTop()
				current++

				wtitle.SetText(fmt.Sprintf(title, current, total))
			} else {
				a.Quit()
			}
		}),
	))

	w.ShowAndRun()
}

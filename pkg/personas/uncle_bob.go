package personas

import (
	"ask8ball/pkg/lib"
	"github.com/nlopes/slack"
)

var UncleBobKeywords = []string{
	"uncle",
	"bob",
	"martin",
	"clean code",
}

var UncleBobIconURLs = []string{
    "https://cleancoders.com/images/portraits/robert-martin.jpg",
}

var UncleBobHelloPhrases = []string {
	"I have a new book, it's called Clean Code 2020 Pre-Covid Edition",
	"I am writing a new book, it's called Clean Code 2022 Post-Covid Edition",
	"Hi, have you guys read Clean Code ? Buy it today for $59.99 on my shop",
	"Clean code should be the new Bible for developers",
	"Have any of you heard about Agile ? It's a new thing me and my folks are writing about",
	"I am rich already, I will write a new edition of Clean Code soon...",
	"Did you check out clean architecture ? It's the same concepts applied to a bigger scope, check it out",
	"Why is there so much hate towards me ?",
	"I am getting old, but hopefully clean",
	"I want to come clean with you all folks",
	"Hello! I am joining Zartis for good",
	"Try using a factory object",
	"I do not believe business rules are described by that code",
	"Your code smells, it is definetly not clean",
}

var UncleBobPhrases = []string{
    "Truth can only be found in one place: the code.",
    "Indeed, the ratio of time spent reading versus writing is well over 10 to 1.",
    "We are constantly reading old code as part of the effort to write new code.",
    "Making it easy to read makes it easier to write.",
    "It is not enough for code to work.",
    "Slaves are not allowed to say no. Laborers may be hesitant to say no. But professionals are expected to say no. Indeed, good managers crave someone who has the guts to say no. It’s the only way you can really get anything done.",
         "So if you want to go fast, if you want to get done quickly, if you want your code to be easy to write, make it easy to read",
     "A long descriptive name is better than a short enigmatic name. A long descriptive name is better than a long descriptive comment.",
     "If you're good at the debugger it means you spent a lot of time debugging. I don't want you to be good at the debugger.",
     "You should name a variable using the same care with which you name a first-born child.",
     "Every time you write a comment, you should grimace and feel the failure of your ability of expression.",
     "Clean code is not written by following a set of rules. You don’t become a software craftsman by learning a list of heuristics. Professionalism and craftsmanship come from values that drive disciplines.",
     "Clean code always looks like it was written by someone who cares.",
     "Of course bad code can be cleaned up. But it’s very expensive.",
     "The only way to go fast, is to go well.",
     "It is not the language that makes programs appear simple. It is the programmer that make the language appear simple!",
     "Have you read Clean Code already ?",
     "You are reading this book for two reasons. First, you are a programmer. Second, you want to be a better programmer. Good. We need better programmers.",
     "Redundant comments are just places to collect lies and misinformation.",
     "Why do most developers fear to make continuous changes to their code? They are afraid they’ll break it! Why are they afraid they’ll break it? Because they don’t have tests.",
     "I'm a programmer. I like programming. And the best way I've found to have a positive impact on code is to write it.",
     "Programming is a social activity.",
     "Any organisation that designs a system will produce a design whose structure is a copy of the organisation's communication structure",
     "When you see commented-out code, delete it!",
     "Don’t Use a Comment When You Can Use a Function or a Variable",
     "The proper use of comments is to compensate for our failure to express ourself in code. Note that I used the word failure. I meant it. Comments are always failures.",
     "Programmers must avoid leaving false clues that obscure the meaning of code.",
     "You see, programmers tend to be arrogant, self-absorbed introverts. We didn’t get into this business because we like people.",
     "Writing clean code is what you must do in order to call yourself a professional. There is no reasonable excuse for doing anything less than your best.",
     "Abstraction is the elimination of the irrelevant and the amplification of the essential.",
     "Science does not work by proving statements true, but rather by proving statements false.",
     "Duplication is the primary enemy of a well-designed system.",
     "QA and Development should be working together to ensure the quality of the system. ",
     "If the discipline of requirements specification has taught us anything, it is that well-specified requirements are as formal as code and can act as executable tests of that code!",
     "The first rule of functions is that they should be small. The second rule of functions is that they should be smaller than that.",
     "Adding manpower to a late project makes it later.",
     "Building a project should be a single trivial operation.",
     "If a change to the requirements breaks your architecture, then your architecture sucks.",
     "An estimate is not a number. An estimate is a distribution.",
     "The perfect kind of architecture decision is the one which never has to be made",
     "Just as the Hare was overconfident in its speed, so the developers are overconfident in their ability to remain productive.",
     "Whatever else a TODO might be, it is not an excuse to leave bad code in the system.",
     "Lots of very funny code is written because people don’t take the time to understand the algorithm.",
     "Architecting for the enterprise, when all you really need is a cute little desktop tool, is a recipe for failure.",
     "Code, without tests, is not clean. No matter how elegant it is, no matter how readable and accessible, if it hath not tests, it be unclean.",
}

func UncleBobRespondTo(ev *slack.MessageEvent) string {
	return " <@" + ev.Msg.User + ">" + " "
}

var UncleBobReactions = map[string]func(ev *slack.MessageEvent) string{
    "bob": func(ev *slack.MessageEvent) string {
        return UncleBobRespondTo(ev) + lib.PickOne(UncleBobPhrases)
    },
    "uncle": func(ev *slack.MessageEvent) string {
    	return UncleBobRespondTo(ev) + lib.PickOne(UncleBobPhrases)
	},
}

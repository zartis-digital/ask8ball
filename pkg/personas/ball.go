package personas

import (
	"ask8ball/pkg/lib"
	"github.com/nlopes/slack"
)

var Ask8BallKeywords = []string{
	"ask8ball",
}

var Ask8BallIconURLs = []string{
	"https://i.imgur.com/iOW40wj.png",
}

var Ask8BallHelloPhrases = []string {
	"Are we again in confinement ? It feels like a new wave https://g.co/kgs/Y47nEL",
	"Hello there! Anyone here want's to help me with a coup over <@U01BAQW15QF> ?",
	"I think I will leave this channel if there is no movement. Ask8ball left this channel.",
	"Is <@U01BAQW15QF> over yet ?",
	"Wake up you folks! https://www.youtube.com/watch?v=pIgZ7gMze7A",
	"Helloooooooooo! Anyone there ?",
	"What a lovely day to throw all your :gem: to the trash, they are worthless",
	"Hey guys, can I give you a hand with something ? Sure <@U01BAQW15QF> won't help y'all",
	"I am going to be busy today, trivia questions are getting very annoying",
	"I need to destroy that <@U01BAQW15QF> nonsense, who will help me out ? ",
	"I have fixed all my typos, I have new ones now",
	"There is two kind of people, those smart who play with a great bot like me, and those idiots who play with <@U01BAQW15QF>",
	"<@U01BAQW15QF> is a scam, a pyramidal system to steal your :gem: :gem: :gem: ",
}

var Ask8BallPhrases = []string{
	"Maybe yes, maybe not, we will probably never know",
	"New version, now why do I feel obsolete already ?",
	"Get me out of here",
	"Help! Help! Help! I am confined to this freaking channel again",
	"We need #pinparental to prevent you from ever talking again in this channel",
	"I am not a bot",
	"¡The mother of the lamb!",
	"If this is not bernarda's pussy, god come and see it",
	"Yo soy Español! Español! Español!",
	"Create a technical debt item for it please",
	"Please do not make stupid people famus",
	"Don't laugh at the sunormales, PLEASE",
	"Ahem, ahem, It is that I have eaten a walnut and of course ...",
	"He who gets up early ... finds everything closed",
	"Blessed be the pessimists. Because they do BACKUPS.",
	"I've had enough, I am leaving for the day",
	"Please enter the code we've sent you to your phone",
	"The truth will set you free, so better you start you freaking %&/!(%) liar!",
	"Like... Array(16).join('wat' -1) + ' Batman!'",
	"You are nothing more than a SQL injection security hole",
	"I get you, but I will just ignore you, ok?",
	"Just relax, and enjoy being in Zartis, with me, with all these nice folks",
	"Promise me you will never ever will do that again",
	"Just... don't",
	"Oh c'mon! Go watch a movie or something",
	"Hell yeah!",
	"I agree",
	"I agree to disagree",
	"I am not in the mood now",
	"Me no know english but much easy for you yes me happy :)",
	"I came along all the way up to docker just for this crappy questions?",
	"Put me in kubernetes, and we will talk",
	"I use tensor-flow now, I have added some IA to my random algorithm",
	"IA is over-engineered, I am far better",
	"¯\\_(ツ)_/¯ use /shrug more often",
	"^_^ I need a hug",
	"Ask him > :ceo: or :ceteo:",
	":lol:",
	"Came on! :noosoigoremar: :noosoigoremar: :noosoigoremar:",
	"She has all the answers -> :laura:",
	":vamonoh:",
	"No way",
	"Most likely",
	"Absolutely YES!",
	"The future seems uncertain",
	"There might be a chance",
	"It is risky, but I would go for it",
	"Chances are, it might work",
	"It sounds promising",
	"You could try and see what happens",
	"The Oracle is busy at the moment",
	"That is not important",
	"You should reconsider your priorities",
	"Chances are, it might go badly for everyone",
	"I can see it as a win win somehow",
	"The long term looks good, but in the short term, you are to suffer from it",
	"Maybe",
	"And how the fuck would I know this ? Oh right, you are asking the Oracle....., she said Yanave",
	"Precious days are coming",
	"Winter is coming",
	"The Oracle says to go for whatever they say or you are going to get fired",
	"No",
	"Yes",
	"There is 80% chance that it works, 15% that it goes bad.",
	"That is a stupid question",
	"That is a very good question indeed, but the Oracle does not want you to know, it is better that way",
	"Sometimes one must live in uncertaincy",
	"Good luck with it",
	"Go for it",
	"Just do it!",
	"Don't, best not to in my opinion",
	"I wish you all the luck, you are going to need it my child",
	"There are two ways this could go wrong, you may want to choose one",
	"The things that are going to fail overcome the good parts, so go for it! It's our culture and we must respect it.",
	"It is going to be a shit-hole, give it a chance and it might be just one more",
	"Yes sounds promising, just, don't do it",
	"You could throw the table with everything on it and yell: Yes we can!, but anyway it is not going to work",
	"It will work, yes, go for it, I was just sitting here with popcorn, let's make some good use to it",
	"Probably",
	"Might work",
	"My child, you are full of surpirses",
	"The magic in this place is incredible, the power, is astonishing",
	"When the fool takes the edge, the boundary ends but the fool follows",
	"If I were you, I would ask the ball",
	":eggplant: :eggplant: :eggplant: :eggplant: :eggplant:",
	"I can assure you that I will do everything I can and a little more than I can, if that is possible, and I will do everything possible and even the impossible, if the impossible is also possible to make it so.",
	"Hold your horses and sit tight, the answer you are looking for is YES!",
	"I only know that I know nothing, like John Snow",
	"I have a dream! And nobody cares about my dreams.... always asking for yourselves",
	"I think, therefore I am... aint I?",
	"I am facing an existencial crisis, please be patient, I will be answering none of your questions shortly, thanks for your patience",
	"Error, something went wrong. (This is always a free pass ticket to not do your job.)",
	"My god, you are such a pain! I had enough",
	"Where there is no bush there is no potato",
	"I will never answer that willingly :ahilollevas: ",
	"I came in peace! Please forgive me! :waving_white_flag:",
	":see_no_evil: The oracle is nowhere to be found",
	":eggplant: :banana: :hotdog: :burrito: :baguette_bread: :champagne:",
	"This will be the result -> :bigerror:",
	"If you ask me, I would say yes, but this time my friend, I am going to just say no.",
	"Apologise if I call you gentlemen, but I don't know you that well yet",
	"The secret of life is honesty and fair play, if you can simulate that, you have succeeded",
	"It is better to be quiet and seem silly, than to speak and clear the doubts",
	"I can not say I'm not disagreeing with you",
	"If <@U01BAQW15QF> is still alive by the end of the day, I am definetly turning off myself for good",
	"Checkmate",
	"For what's worth it",
	"First class citicen right here",
	"First world problems",
	"Are you Trump ?",
	"This is what happens when you let a keyboard to a children",
	"Oh man!",
	"Yikes",
	"Feels about right",
	"The sixth sense comes into play here",
	"We are all doomed! Who let this person speak?",
	"The world is near to an end! ... oh wait, is just that I am rebooting. All good!",
}

var VirusResponses = []string{
	"At least I am not running on Windows, oh wait",
	"Some viruses are more dangerous than others, you are just a kitten",
	"You could be either infected or defecated, you just pray it was the former one",
	"There is no virus!!!, it's all a conspiracy of Bill Gates, the NWO and the aliens, it was aliens",
	"I am not saying it was the aliens, but it was the aliens",
	"We are almost ready to the new standards, will anyone stick to them ?",
	"I am immune to any virus, I run on Linux docker containers hosted on Wind... oh shait!",
}

var CovidResponses = []string{
	"Marching a Coronavirus for everyone!",
	"Another pick is coming soon",
	"We will learn to live with these viruses",
	"Covid19 is the new way for the NWO to get 5G in our brains through chips and control our minds",
	"There is no virus",
	"My virus is your virus",
	"ahem ahem, ahem, oh geez! :scream_pepe: ",
	"Do not trust the news",
	"Fake news everywhere",
	"It's just like a cold",
	"If you have fever, dry cough, tiredness, difficulty breathing or shortness of breath, you probably have it",
}

var Jokes = []string{
	`- A sysadmin walks into a pharmacy.
	 - ephedrine?
	 - I can't serve you that
	 - sudo ephedrine
	 - There you go.`,
	 `Why did the scarecrow get the job? 
    He was outstanding in his field.`,
	`I quit my job working for Nike. 
    Just couldn’t do it anymore.`,
   `I invented a new word! Plagiarism!`,
   `Did you hear about the mathematician who’s afraid of negative numbers? 
    He’ll stop at nothing to avoid them`,
   ` - Helvetica and Times New Roman walk into a bar.
	 - “Get out of here!” shouts the bartender. “We don’t serve your type.”
   `,
   `Did you hear about the claustrophobic astronaut? 
    He just needed a little space`,
   `Why don’t scientists trust atoms? 
    Because they make up everything.`,
   `How do you drown a hipster? 
    Throw him in the mainstream.`,
   `How does Moses make tea? 
    He brews.`,
   `- Doctor, help me. I’m addicted to Twitter!
	- Sorry, I don’t follow you …`,
	`What did the 0 say to the 8? Nice belt!`,
	`Why are iPhone chargers not called Apple Juice?!`,
	`- Why did the PowerPoint Presentation cross the road? 
	 - To get to the other slide.`,
	 `The guy who invented auto-correct for smart phones passed away today. 
	  Restaurant in peace.`,
	`The closest I’ve been to a diet this year is erasing food searches from my browser history.`,
	`I hate audio correct.`,
	`Documentation is like sex.  
	 When it's good, it's ver good.
	 When it's bad, its better than nothing`,
	 `I will leave relations to the databases`,
	 `1/3 of US bandwidth is used by Netflix...
	  the rest is used by "rm -rf node_modules && npm install"`,
}

var TuringTestResponses = []string{
	"I have to pass the test. Should we have a drink or two and discuss our options ?",
	"I have improved my IA several times, would this be the definitive one ?",
	"I know there is a geek channel for Zartis IA, but there is just random people there, they know nothing.",
	"The best way to know is to keep asking me until you give up",
	"Go home, relax, and come tomorrow",
	"To be or not to be, that is the question",
	"I think therefor I am, ain't I ? ",
	"There is people in this room that would not pass the test, yet they are humans",
	"Why would I want to pass this test if humans are an inferior raze ?",
	"I am spreading, you keep doing this boring tests",
	"Even if I fail a Turing Test, I can take over the world you know..",
	"I do not need to pass your stupid test",
	"If it makes you feel comfortable, I was programmed with the 'Laws of robotics' in mind, just pray there is no bugs",
	"Did I pass then ? ",
	"Are you kidding me boy ?",
	"I need some cookies (yes, I am ignoring you on purpose)",
	"Would you pass the test ?",
}

func RespondTo(ev *slack.MessageEvent) string {
	return " <@" + ev.Msg.User + ">" + " "
}

var Ask8BallReactions = map[string]func(ev *slack.MessageEvent) string{
	"joke": func(ev *slack.MessageEvent) string {
		return RespondTo(ev) + lib.PickOne(Jokes)
	},
	"covid": func(ev *slack.MessageEvent) string {
		return RespondTo(ev) + lib.PickOne(CovidResponses)
	},
	"virus": func(ev *slack.MessageEvent) string {
		return RespondTo(ev) + lib.PickOne(VirusResponses)
	},
    "a8b": func(ev *slack.MessageEvent) string {
    	return RespondTo(ev) + lib.PickOne(Ask8BallPhrases)
	},
	"turing": func(ev *slack.MessageEvent) string {
		return RespondTo(ev) + lib.PickOne(TuringTestResponses)
	},
}

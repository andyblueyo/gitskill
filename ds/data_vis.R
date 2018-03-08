library(dplyr)
library(tidyr)
library(ggplot2)

# Read data
repo_data <- read.csv("merged_data.csv", stringsAsFactors = FALSE)

colnames(repo_data)

# Gather data with language and line count
gather_data <- repo_data %>% select(-c(X)) %>% gather(language, line, Java:Zimpl)

pls <- na.omit(gather_data)
count_lang <- gather_data %>% filter(line != 0) %>% count(language) %>% filter(n > 4400)

count_lang <- gather_data %>% select(language, stars, forks, line, publicRepos) %>% filter(line != 0) %>% group_by(language) %>% count(stars)

total_lines <- pls %>% select(language, stars, forks, line, publicRepos) %>% filter(stars != 0) %>% group_by(language) %>% summarise_at(c("lines"), sum)

# bar plot of freq of counts of language
bar_lang_count <- ggplot(data=count_lang, aes(x = language, y = n)) + geom_bar(stat = "identity") + 
  theme(axis.text.x = element_text(angle = 90, hjust = 1, size = 12), plot.title = element_text(hjust = 0.5)) + labs(x = "Languages", y = "Counts", title = "Frequency of Languages")
bar_lang_count

#scatter of stars and forks
forks_star <- ggplot(data = pls, aes(x=forks, y = stars)) + geom_point() +
  theme(axis.text.x = element_text(angle = 90, hjust = 1, size = 12), plot.title = element_text(hjust = 0.5)) + labs(x = "Number of forks", y = "Number of Stars", title = "Relationship with Stars and Forks")
forks_star



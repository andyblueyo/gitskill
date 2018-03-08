library(dplyr)
library(tidyr)
library(ggplot2)

# Read data
repo_data <- read.csv("merged_data.csv", stringsAsFactors = FALSE)

colnames(repo_data)

# Gather data with language and line count
gather_data <- repo_data %>% select(-c(X)) %>% gather(language, line, Java:Zimpl)

pls <- na.omit(gather_data)
count_lang <- pls %>% filter(line != 0) %>% count(language)
 
total_lines <- pls %>% select(language, stars, forks, line, publicRepos) %>% filter(stars != 0) %>% group_by(language) %>% summarise_at(c("stars"), sum)

# bar plot of freq of counts of language
bar_lang_count <- ggplot(data=count_lang, aes(x = language, y = n)) + geom_bar(stat = "identity") + 
  theme(axis.text.x = element_text(angle = 90, hjust = 1, size = 5))

#scatter of stars and forks
forks_star <- ggplot(data = gather_data, aes(x=forks, y = stars)) + geom_point()



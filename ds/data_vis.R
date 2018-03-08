library(dplyr)
library(tidyr)
library(ggplot2)

# Read data
repo_data <- read.csv("merged_data.csv", stringsAsFactors = FALSE)

colnames(repo_data)

# Gather data with language and line count
gather_data <- repo_data %>% select(-c(X)) %>% gather(language, line, Java:Zimpl)


count_lang <- gather_data %>% filter(line != 0) %>%  count(language)

# bar plot of freq of counts of language
ggplot(data=count_lang, aes(x = language, y = n)) + geom_bar(stat = "identity") + 
  theme(axis.text.x = element_text(angle = 90, hjust = 1, size = 5))

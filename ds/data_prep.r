library(dplyr)
library(jsonlite)
library(tidyr)

repo_data <- fromJSON("./data/repos.json")
user_data <- fromJSON("./data/users.json")


user_data <- flatten(user_data, recursive = TRUE)
View(user_data)
repo_data <- flatten(repo_data, recursive = TRUE)
View(repo_data)

print(repo_data[2,6])
df <- data.frame(repo_data[,c("name", "languages", "ownerType", "ownerName")])



z <- 1
for(language in df$languages){
  if(length(language)>0){
    
    lines <- language[,"lines"]
    lang <- language[,"name"]
    if(length(lang)>0){
      
      for(i in 1:length(lang)){
        print(z)
        
        df[z, lang[i]] <- lines[i]
      }
    }
  }
  
  z <- z + 1
}
drops <- c("languages")
df2 <- df[, !(names(df) %in% drops)] #%>% apply(2,as.character)
df2[is.na(df2)] <- 0


df3 <- df2 %>% 
  group_by(ownerName) %>% 
  filter() %>% 
  summarize_all(funs(sum(as.numeric(.))))

write.csv(x = df2, file = "./data/expanded.csv")
write.csv(x = repo_data, file = "./data/repos.csv")

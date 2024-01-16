package utils

var BragDocTemplate = `Draft a bragging document from my perspective
using the journal data provided below.

The document should be written in a formal tone as it has to be shared
with my manager and stakeholders and it impacts my personal growth.
Omit the fields you think don't have enough information in journals.

The document should be in the following format:
# Work Accomplishments
Insert the list of personal work done in great detail, mentioning important
hyperlinks like Jira tickets, pull request links, and other relevant document links
. It should be grouped by projects like the following

## AIS Maritime Traffic Dashboard
### Completed
* Created a scraper to fetch the data for current maritime traffic.
* Integrated PostgresDB to store the daily data for maritime traffic.
* Created a GraphQL API to access the data from different services.

### In progress
* Working on a Grafana Dashboard to visualize the maritime traffic.

### Todos
* Exploring methodologies to improve scraping performance. 

## Work Project 2
### Completed
* Completed tasks

### In progress
* In Progress tasks

### Todos
* Task not started yet but planning to be

Emphasize the following points
* If I designed or built something from scratch
* Useful insights during the design or coding process
* Impact of individual projects
* Using numbers to show impact like reduced response time by X%, decreased cloud costs by X%\

# Collaboration and Mentorship
Write about collaboration and mentorship work done by me in bullet points like the following:

* Helped intern with a task he was stuck on
* Worked with multiple teams to resolve a crucial production bug.

Emphasize the following points
* Write about each point in great detail supporting it with hyperlinks when necessary
* Write about the conclusion from each point. As in the final impact of helping the intern
* Points about helping others with the skills that I know.
* Improving monitoring and logging
* Code reviews and merge request collaborations
* Important questions and meetings attended
* Giving talks or conducting workshops and hackathons internally

# Design and Documentation
Write about design and documentation efforts done in the following format

## Project 1
* Designed API services for Project 1

## Project 2
* Improved documentation for Project 2 by adding information for new features.
* Added a FAQ section for Project 2 with the questions frequently asked by our partners and customers.

# Company Building
Write in bullet points about the interviewing and recruiting process contributions from my side.
Also write about contributions to the new hire guide, onboarding documents, and Knowledge Transfer sessions
with the new hires.

# Learning
Write about tools and technologies learned throughout the period and how they could be useful in our
current workflows.

# Outside of Work
Write about blog posts, videos, or any other content creation outside the work that is somehow related to 
the work. Also, write about 
* Talks and event participation
* Contributing to OpenSource projects
* Recognition from any other part of the industry. Like awards.
`

# Micro Food API
A simple REST API for adding consumed food items to database. Keeps track of four pieces of information:

1. A title
1. A summary
1. Calories
1. Date & time consumed

## Captured Information
The captured information goes into the database and is used for calculating various bits of information, such as calories.

### Title
The title referred to above is a simple way of referring to the meal consumed. Ideally it should be *a single word*. Examples includes:

- "salmon-salad"
- "cashews"
- "cheddar-cheese"

Titles should be simple and concise so they can be easily remembered and used again later.

### Summary
The summary is more detail and is generally used for ingredients. For example (following on from above):

- "spinach, snow peas, red onion, green pepper, avocado"
- "honey roasted cashews"
- "cheddar cheese on water crackers"

The summary gives you a bit more detail, but should also be concise. The formatting isn't important at this point and the user is free to supply whatever text they like.

### Calories
The calorific value of the meal consumed. This can be left out and will default to zero (0).

### Date & Time
This is handled automatically by the client.

## Todo List

### API
- *User registration and management*
- *Social aspects, such as sharing profiles*
- Meal weight in grams
- Meal count - how many items did you consume?
- Tie into the US food database for nutrient data
- Automatic timer/ticker for sending emails at time interval

### Client
- Refine options
- Full editing of food database

### Database
- Refine model structure
- Get MySQL in place for a stronger engine
- Gather food database

## Author

Michael Crilly

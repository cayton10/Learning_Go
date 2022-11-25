LEARN GO
# Comic Mischief
Use variables and types in the Go programming language to maintain the catalog of items in your comic book store Comic Mischief.

Tasks
0/23 Complete
Mark the tasks as complete by checking them off
Creating Our Basic Go Program
1.
Let’s start by making a basic Go program. Before anything, we should declare our package information. In Go, this is done with the line package main. Type this as the first line of the program.


Stuck? Get a hint
2.
Next we want to print out our data to the terminal, so we’ll be using the fmt package. Import fmt into our Go program.


Stuck? Get a hint
3.
Now let’s create our base program by adding in a main function.


Stuck? Get a hint
Defining Our Comic Book Variables
4.
Now let’s define the variables we’re going to use to store the data for each comic book. Comics have: a publisher, a writer, an artist, a title, a year they were released in, a grade for its condition, and a page number.

Let’s start by defining our string-type data. Define variables for publisher, writer, artist, and title. Make all of these variables strings!


Stuck? Get a hint
5.
Let’s start defining the numeric variables, we have two integers that we need to define: the year and the page number.

Define integer variables called year and pageNumber.


Stuck? Get a hint
6.
We’ll need to define our condition grade. This is a score from 0.0 to 10.0 with decimal values. Because of the decimal values, we’ll need to use a float to store it.

Create a grade variable that is a float32.


Stuck? Get a hint
7.
Let’s use all of our variables in a series of fmt.Println statements! We want to output all the information from each of our comic books in a way that is repeatable, so be sure to format all of our information in a human-readable way.

We can use commas to separate string literals from variables and Println() will give us spaces between the things we input, so we can write something like:

fmt.Println(title, "written by", writer, "drawn by", artist)
Fill in the rest of the variables to give a full depiction of the comic book in question!


Stuck? Get a hint
Assigning Our Comic Book Variables: Mr. GoToSleep
8.
We finally have our first comic book to keep track of, let’s use software to manage it!

We’ll be adding this variable and the variables for the following steps above the print statement we previously just created. The honor of our first comic book goes to the award-winning "Mr. GoToSleep". Assign that value to title.

Click the save button to see the output change in the terminal. As we continue to assign variables their values, save the code and watch the outputs change as well!


Stuck? Get a hint
9.
We need to give the author her due credit, store the author’s name "Tracey Hatchet" in our writer variable.

10.
Let’s keep the momentum going by storing the artist, the American legend herself "Jewel Tampson".

11.
Be sure to save the publisher of this graphic novel, "DizzyBooks Publishing Inc." as well.

12.
This copy of Mr. GoToSleep came out in 1997 so let’s use that value as our year.

13.
It’s a floppy comic so it’s only got 14 pages. Save that value in our pageNumber variable.

14.
Our beloved vintage Mr. GoToSleep comic has unfortunately seen some wear-and-tear in the intervening decades and is currently valued at a 6.5 in terms of condition. Assign that value to its grade.

Remember to save the code to see the full message with their intended values printed!

Reassigning our Comic Book Variables: Epic
15.
We’re going to add another comic book, but since we want to keep the print out of first comic’s details, copy the existing print statement and paste the new statement right below existing statement.

We’ll be assigning new variables in between these two print statements to observe how values change depending on where in the code they’re located.


Stuck? Get a hint
16.
Our next comic book is the esteemed first trade paperback of the incredible comic Epic. Reassign title to "Epic Vol. 1", remember add this code in between the two print statements.

Save the code to see how our output changes!


Stuck? Get a hint
17.
The author, a noted comics writer in his own right, is named Ryan N. Shawn. Reassign the writer variable to give the writer his due credit.


Stuck? Get a hint
18.
Reassign artist with our current artist’s name, the Spiegeler-winning "Phoebe Paperclips".

19.
We have the same publisher as last time "DizzyBooks Publishing Inc.", so we don’t have to reassign it here. But we can if we were trying to be diligent about having all the information in one place.


Stuck? Get a hint
20.
Epic Vol. 1 was released in 2013, so use that number to update our year variable.

21.
As a trade paperback it’s a compilation of a few floppy comics and comes out to 160 pages. Use that value to store over our pageNumber variable.

22.
This comic is in fantastic condition but has a small shipping defect, warranting a 9.0 for its grade.

Finished!
23.
Good job, you’ve gotten plenty of experience creating new variables and you’re set up to learn what’s next!

If you want to challenge yourself, consider:

Adding more variables that store a comic’s data (like the genre).
Create another comic series, with own unique set of variables.
Perform some computation using the age, grade, and pageNumber to calculate the cost of the comic book.
You can do it!
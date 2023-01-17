- organizin data to appear similar accross all records and fields
- organizing data into tables so data organized the same across
- cut down redunancy
  - removing duplicates creates more space
- make database more efficient
- organizing data into database
  - create tables
  - establish relationship between tables
- provides logic and structure to data

Normal Forms
-------------------------
First Normal Form (1NF)
- no repeated data
- single value for cell for data entry
- each record uniq

Second
- follows 1NF
- single primary key
- subsets of data like the type of cookie should be moved to a dif table referenced by FK

Third
- follows 2NF
- data must only be dependant on the primary key
  when primary key is changed all impacted data placed in new table
  ex: changing name might also change gender so gender placed in a new table.

Advantages of Normalization :
Here we can perceive any reason why Normalization is an alluring possibility in RDBMS ideas.
- A more modest information base can be kept up as standardization disposes of the copy information. Generally speaking size of the information base is diminished thus.
- Better execution is guaranteed which can be connected to the above point. As information bases become lesser in size, the goes through the information turns out to be quicker and more limited in this way improving reaction time and speed.
- Narrower tables are conceivable as standardized tables will be tweaked and will have lesser segments which considers more information records per page.
- Fewer files per table guarantees quicker support assignments (file modifies).
A- lso understands the choice of joining just the tables that are required.

Disadvantages of Normalization :
- More tables to join as by spreading out information into more tables, the need to join table’s increments and the undertaking turns out to be more dreary. The information base gets more enthusiastically to acknowledge too.
- Tables will contain codes as opposed to genuine information as the rehashed information will be put away as lines of codes instead of the genuine information. Thusly, there is consistently a need to go to the query table.
- Data model turns out to be incredibly hard to question against as the information model is advanced for applications, not for impromptu questioning. (Impromptu question is an inquiry that can’t be resolved before the issuance of the question. It comprises of a SQL that is developed progressively and is typically built by work area cordial question devices.). Subsequently it is difficult to display the information base without understanding what the client wants.
- As the typical structure type advances, the exhibition turns out to be increasingly slow.
- Proper information is needed on the different ordinary structures to execute the standardization cycle effectively. Reckless use may prompt awful plan loaded up with significant peculiarities and information irregularity.
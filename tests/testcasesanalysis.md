##  **Golden Test Set (Functional test cases — natural language)**

Παρακάτω σημειώνω βασικά functional tests (όσα είναι στα audit examples) και μετά πρωτότυπα, tricky παραδείγματα

### **A. Βασικά / audit παραδείγματα**

**cap και up**

Input: `it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom.`

output: `It was the best of times, it was the worst of TIMES, it was the age of wisdom.`

**Hex και  Bin**

Input: `Simply add 42 (hex) and 10 (bin) and you will see the result is 68.`

output: `Simply add 66 and 2 and you will see the result is 68.`

**καμία αλλαγή**

Input: `There is no greater agony than bearing a untold story inside you.`

output: ίδιο κείμενο χωρίς αλλαγές.

**σημεία στίξης κενά**

Input: `Punctuation tests are ... kinda boring ,what do you think ?`

output: `Punctuation tests are... kinda boring, what do you think?`

**Απόστροφοι**

Input: `I am exactly how they describe me: ' awesome '`

output: `I am exactly how they describe me: 'awesome'`

###  **Πέντε πρωτότυπα, tricky παραδείγματα (να καλύψουν edge cases)**

**`(up, N)` εφαρμόζεται και σε λέξεις που βρίσκονται πριν ή μετά από σημεία στίξης**

Input: `we saw the bright, and shining sun. (up, 4) .`

Σημείο: οι 4 προηγούμενες λέξεις μπορεί να περιλαμβάνουν κόμμα ή τελεία — πρέπει να μετρήσει μόνο πραγματικές λέξεις και να τους αλλάξει το case.

output: `we saw the BRIGHT, AND SHINING SUN.`

**Quotes** 

Αν μια εντολή όπως `(up)`, `(low)`, `(cap)` εμφανιστεί μέσα σε quotes (' '),  
   
 **Input: `She said: 'I will shout (up) now' and left.`**  
 **Output: `She said: 'I will SHOUT now' and left.`**

**πολλαπλές εντολες σε λιγο κειμενο τι θα εφαρμοσει πρώτα**

Input: `test 1E (hex) (up) sample`

Σημείο: Πρέπει να αποφασίσουμε αν κάνουμε πρώτα hex → 30 μετά up → `30` Output: `test 30 sample` →

**a → an with H και κεφαλαία**

Input: `She saw a Honest man. She found a house.`

Σημείο: λέξεις που αρχίζουν με h: είτε `a honest` → `an honest` 

output: `She saw an Honest man. She found a house.` 

**Πολλά σημεία στίξης**

Input: `Wait ! ? Why... no ?`

Σημείο: `! ?` πρέπει να γίνουν `!?` ή `!?` συνδεδεμένα; `...` πρέπει να γίνει συνέχεια. Επίσης καθαρισμός κενών.

Output: `Wait!? Why... no?`

it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom.   
Simply add 42 (hex) and 10 (bin) and you will see the result is 68\.   
There is no greater agony than bearing a untold story inside you.   
Punctuation tests are ... kinda boring ,what do you think ?   
I am exactly how they describe me: ' awesome '.   
we saw the bright, and shining sun. (up, 4\) .   
She said: 'I will shout (up) now' and left.   
test 1E (hex) (up) sample.   
She saw a Honest man. She found a house.   
Wait \! ? Why... no ?   
it (cap) was a incredible day , the sun was shining , and we saw 1E (hex) birds flying across the sky (up) \!   
then , john said : '  this is a happy moment  ' , but i told him it was 101 (bin) times better (cap,5) .   
later , a honest man approached and whispered : '  YOU ARE A HERO  ' (low,4) \!\!   
in the distance , we could hear an echo of laughter ... it felt like a dream (cap,5) , truly unforgettable (up,2) \!   
and finally , we all shouted victory (up) together , because it was ours .

It was the best of times, it was the worst of TIMES, it was the age of wisdom.   
Simply add 66 and 2 and you will see the result is 68\.   
There is no greater agony than bearing an untold story inside you.   
Punctuation tests are... kinda boring, what do you think?   
I am exactly how they describe me: 'awesome'.   
we saw the BRIGHT, AND SHINING SUN.   
She said: 'I will SHOUT now' and left.   
test 30 sample.   
She saw an Honest man. She found a house.   
Wait\!? Why... no?   
it was an incredible day, the sun was shining, and we saw 30 birds flying across the SKY\!   
then, John said: 'This is a happy moment', but I told him It Was 5 Times Better.   
later, an honest man approached and whispered: 'you are a hero'\!\!   
in the distance, we could hear an echo of laughter... It Felt Like A Dream, TRULY UNFORGETTABLE\!   
and finally, we all shouted VICTORY together, because it was ours.


# emission(1) - do calculations on measurements from the author's valve tester

Kozmix Go, 17-FEB-2023

```
emission <valve type> ht=... curr=... [nb=...]
```


<a name="description"></a>

# Description

**emission**
processes the test results from the author's valve tester. The command
line arguments are:


* **designation**  
  The valve type (EL84, ECC83, etc; case-insensitive).
* **ht=...**  
  HT (anode) voltage as measured from the red socket, in volts. This
  will be around 150V for preamp valves and around 300V for power
  valves.
* **curr=...**  
  Cathode current as measured from the green socket (where 1mV measured
  equals 1mA of current).
* **nb=...**  
  Negative bias voltage as measured from the blue socket and set with
  the knob to the left of it (optional).
  

**emission**
will then calculate power dissipation (both at 100% for cathode biased
valves and at 70% for fixed biased valve), and emit warnings if any
measured values exceed maximum values (or in case of negative bias,
are more than 50% off the recommended testing value).


<a name="examples"></a>

# Examples


.EX
$ emission el84 curr=48m ht=297 nb=14
valve: EL84

value                        measured  nominal/max.
anode voltage            V =    297 V    300 V
negative bias            V =    -14 V  -13.5 V
cathode current          I =    48 mA    65 mA
dissipation cathode bias P =  14.25 W     12 W  WARNING
dissipation fixed bias   P =  14.25 W    8.4 W  WARNING

.EE



<a name="see-also"></a>

# See Also

**valve**(1)


<a name="bugs"></a>

# Bugs

**emission**
only has data on a small range of valve types built in, and only knows
them by their European designations.


<a name="author"></a>

# Author

svm


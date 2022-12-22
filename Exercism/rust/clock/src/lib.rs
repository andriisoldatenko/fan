use std::fmt;

pub struct Clock {
    hours: i32,
    minutes: i32,
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        let mut new_minutes: i32 = minutes;
        let mut new_hours: i32 = hours;

        if new_minutes < 0 {
            new_minutes = 60 + minutes;
            new_hours = hours - ((new_minutes / 60) + 1);
            if new_hours < 0 {
                new_hours = 24 + new_hours
            }
        } else {
            new_minutes = minutes % 60;
            new_hours = hours % 24;
        }

        Clock {
            minutes: new_minutes,
            hours: new_hours,
        }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {

        let mut new_minutes: i32;
        let mut new_hours: i32;

        new_minutes = self.minutes + minutes;

        if new_minutes < 0 {
            new_minutes = 60*(new_minutes.abs() / 60 + 1) + new_minutes;
            new_hours = self.hours - ((self.minutes + minutes).abs() / 60 + 1);
            if new_hours < 0 {
                new_hours = 24 + new_hours
            }
        } else {
            new_minutes = (self.minutes + minutes) % 60;
            new_hours = self.hours + (self.minutes + minutes) / 60;
        }

        Clock {
            minutes: new_minutes,
            hours: new_hours,
        }
    }

    pub fn to_string(&self) -> String {
        format!("{self}")
    }
}


impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{:0>2}:{:0>2}", self.hours, self.minutes)
    }
}

impl fmt::Debug for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        f.debug_struct("Clock")
            .field("hours", &self.hours)
            .field("minutes", &self.minutes)
            .finish()
    }
}

impl PartialEq for Clock {
    fn eq(&self, other: &Self) -> bool {
        if self.hours >= other.hours && self.minutes >= other.minutes {
            return true;
        }
        false
    }
}

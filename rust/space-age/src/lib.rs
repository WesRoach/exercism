// The code below is a stub. Just enough to satisfy the compiler.
// In order to pass the tests you can add-to or change any of this code.
use std::time::Duration as StdDuration;

#[derive(Debug)]
pub struct Duration {
    std_duration: StdDuration,
}

impl From<u64> for Duration {
    fn from(s: u64) -> Self {
        return Duration {
            std_duration: StdDuration::from_secs(s),
        };
    }
}

pub trait Planet {
    fn years_during(d: &Duration) -> f64 {
        todo!("convert a duration ({d:?}) to the number of years on this planet for that duration");
    }
}

pub struct Mercury;
pub struct Venus;
pub struct Earth;

pub struct Mars;
pub struct Jupiter;
pub struct Saturn;
pub struct Uranus;
pub struct Neptune;

impl Planet for Mercury {
    fn years_during(d: &Duration) -> f64 {
        let seconds = d.std_duration.as_secs_f64();
        return seconds / 7600543.81992;
    }
}
impl Planet for Venus {
    fn years_during(d: &Duration) -> f64 {
        let seconds = d.std_duration.as_secs_f64();
        return seconds / 19414149.052176;
    }
}
impl Planet for Earth {
    fn years_during(d: &Duration) -> f64 {
        let seconds = d.std_duration.as_secs_f64();
        return seconds / 31557600.0;
    }
}
impl Planet for Mars {
    fn years_during(d: &Duration) -> f64 {
        let seconds = d.std_duration.as_secs_f64();
        return seconds / 59354032.69008;
    }
}
impl Planet for Jupiter {
    fn years_during(d: &Duration) -> f64 {
        let seconds = d.std_duration.as_secs_f64();
        return seconds / 374355659.124;
    }
}
impl Planet for Saturn {
    fn years_during(d: &Duration) -> f64 {
        let seconds = d.std_duration.as_secs_f64();
        return seconds / 929292362.8848;
    }
}
impl Planet for Uranus {
    fn years_during(d: &Duration) -> f64 {
        let seconds = d.std_duration.as_secs_f64();
        return seconds / 2651370019.3296;
    }
}
impl Planet for Neptune {
    fn years_during(d: &Duration) -> f64 {
        let seconds = d.std_duration.as_secs_f64();
        return seconds / 5200418560.032;
    }
}

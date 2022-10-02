# Future Appointment Schedulling

Golang APIs  that allows client to make an appointment with their trainer

I built this service assuming that trainer has already added his/her availability to DB, meaning that it's guaranteed
we have the most up-to-date available/booked appointments at the time of incoming request.

Areas of improvement:
- logging
- openAPI
- configuration files that hold service configs e.g. DB connection string or service host
- routing middleware that checks if url `trainer_id` is valid
- separate appointment validation logic into internal "business_logic" pkg
- API that allows trainers to create available time slots for the week
- public UUIDs for trainers, customers, and appointments

## Usage

### Run service locally

```sh
make run-service
```

### Endpoints

#### GET ```http://localhost:8080/appointments/{trainerID}/available```

Get a list of available appointment times for a trainer between two dates.

#### POST ```http://localhost:8080/appointments/{trainerId}```

Book a new appointment

#### GET ```http://localhost:8080/appointments/{trainerID}/scheduled```

Get a list of scheduled appointments for a trainer.

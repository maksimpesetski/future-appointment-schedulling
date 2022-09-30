
-- create trainer
INSERT INTO trainers (trainer_id, first_name, last_name)
VALUES (1, 'John', 'Strong');

-- create customer
INSERT INTO customers (customer_id, first_name, last_name)
VALUES (2, 'Valuable', 'Customer');

-- create appointment
INSERT INTO appointments (starts_at, ends_at, name, trainer_id, customer_id)
VALUES ('2022-10-25T09:00:00-07:00', '2022-10-25T09:30:00-07:00', 'Fancy workout', 1, 2),
       ('2022-10-25T09:30:00-07:00', '2022-10-25T10:00:00-07:00', 'HITT', 1, null),
       ('2022-10-25T10:00:00-07:00', '2022-10-25T10:30:00-07:00', 'Yoga', 1, null),
       ('2022-10-25T10:30:00-07:00', '2022-10-25T11:00:00-07:00', 'Zumba', 1, null),
       ('2022-10-25T11:30:00-07:00', '2022-10-25T12:00:00-07:00', 'Heavy lifting', 1, null);

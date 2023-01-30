CREATE TABLE "user" (
                            "id" bigserial PRIMARY KEY,
                            "created_at" timestamp NOT NULL DEFAULT NOW(),
                            "email" varchar(255) NOT NULL unique ,
                            "phone" varchar(25) NOT NULL UNIQUE ,
                            "hashed_password" varchar(250) NOT NULL
);
CREATE TABLE user (
                       id UUID PRIMARY KEY,
                       category VARCHAR(50),
                       level VARCHAR(50),
                       verified BOOLEAN,
                       verified_code VARCHAR(50),
                       created_at TIMESTAMP,
                       updated_at TIMESTAMP
);

CREATE TABLE user_token (
                             id SERIAL PRIMARY KEY,
                             user_id UUID REFERENCES user(id),
                             access_token TEXT NOT NULL,
                             created_at TIMESTAMPTZ NOT NULL,
                             updated_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE user_profile (
                               user_id UUID PRIMARY KEY,
                               first_name TEXT NOT NULL,
                               last_name TEXT NOT NULL,
                               phone_number TEXT NOT NULL,
                               email TEXT NOT NULL,
                               country_id INT NOT NULL,
                               created_at TIMESTAMPTZ NOT NULL,
                               updated_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE restaurant (
                            id UUID PRIMARY KEY,
                            name VARCHAR(255) NOT NULL,
                            restaurant_type_id INT REFERENCES restaurant_type(id),
                            description TEXT,
                            address VARCHAR(255),
                            phone_number VARCHAR(20),
                            email VARCHAR(255),
                            website VARCHAR(255),
                            belong_to_user VARCHAR(255),
                            created_at TIMESTAMP,
                            updated_at TIMESTAMP,
                            deleted_at TIMESTAMP
);

CREATE TABLE restaurant_menu (
                                id SERIAL PRIMARY KEY,
                                restaurant_id UUID REFERENCES restaurant(id),
                                name VARCHAR(255) NOT NULL,
                                description TEXT,
                                menu_category_id INT REFERENCES menu_category(id),
                                price VARCHAR(50),
                                currency VARCHAR(10),
                                created_at TIMESTAMP,
                                updated_at TIMESTAMP,
                                deleted_at TIMESTAMP
);

CREATE TABLE menu_category (
                              id SERIAL PRIMARY KEY,
                              name VARCHAR(255) NOT NULL,
                              description TEXT,
                              created_at TIMESTAMP,
                              updated_at TIMESTAMP
);

CREATE TABLE restaurant_type (
                                id SERIAL PRIMARY KEY,
                                name VARCHAR(255) NOT NULL
);

CREATE TABLE restaurant_tag (
                               id SERIAL PRIMARY KEY,
                               name VARCHAR(255) NOT NULL,
                               created_at TIMESTAMP,
                               updated_at TIMESTAMP,
                               deleted_at TIMESTAMP
);

CREATE TABLE restaurant_tag_mapping (
                                      id SERIAL PRIMARY KEY,
                                      restaurant_id UUID REFERENCES restaurant(id),
                                      restaurant_tag_id INT REFERENCES restaurant_tag(id),
                                      created_at TIMESTAMP,
                                      updated_at TIMESTAMP,
                                      deleted_at TIMESTAMP
);


CREATE TABLE restaurant_rating (
                                  id SERIAL PRIMARY KEY,
                                  restaurant_id UUID REFERENCES restaurant(id),
                                  user_id UUID REFERENCES user(id),
                                  overview_rating FLOAT NOT NULL,
                                  cp_rating FLOAT NOT NULL,
                                  meal_rating FLOAT NOT NULL,
                                  drink_rating FLOAT NOT NULL,
                                  service_rating FLOAT NOT NULL,
                                  ambience_rating FLOAT NOT NULL,
                                  comment TEXT,
                                  created_at TIMESTAMP NOT NULL,
                                  updated_at TIMESTAMP,
                                  deleted_at TIMESTAMP
);

CREATE TABLE restaurant_rating_history (
                                         id SERIAL PRIMARY KEY,
                                         restaurant_id UUID REFERENCES restaurant(id),
                                         month VARCHAR(10) NOT NULL,
                                         year VARCHAR(10) NOT NULL,
                                         overview_rating_avg FLOAT NOT NULL,
                                         overview_rating_count INT NOT NULL,
                                         cp_rating_avg FLOAT NOT NULL,
                                         cp_rating_count INT NOT NULL,
                                         meal_rating_avg FLOAT NOT NULL,
                                         meal_rating_count INT NOT NULL,
                                         drink_rating_avg FLOAT NOT NULL,
                                         drink_rating_count INT NOT NULL,
                                         service_rating_avg FLOAT NOT NULL,
                                         service_rating_count INT NOT NULL,
                                         ambience_rating_avg FLOAT NOT NULL,
                                         ambience_rating_count INT NOT NULL,
                                         created_at TIMESTAMP NOT NULL,
                                         updated_at TIMESTAMP,
                                         deleted_at TIMESTAMP
);


CREATE TABLE organization (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    belong_to_user UUID REFERENCES user(id)
);

CREATE TABLE organization_member (
    organization_id UUID REFERENCES organization(id),
    user_id UUID REFERENCES user(id),
    role_id VARCHAR(50),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (organization_id, user_id)
);

CREATE TABLE organization_role (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    category VARCHAR(50),
    belong_to_organization UUID REFERENCES organization(id),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE fundraising (
                             id UUID PRIMARY KEY,
                             user_id UUID REFERENCES user(id),
                             restaurant_id UUID REFERENCES restaurant(id),
                             title VARCHAR(255) NOT NULL,
                             status VARCHAR(50) NOT NULL,
                             lowest_target_price FLOAT NOT NULL,
                             success_target_price FLOAT NOT NULL,
                             currency VARCHAR(10) NOT NULL,
                             introduction TEXT,
                             end_date TIMESTAMP,
                             created_at TIMESTAMP,
                             updated_at TIMESTAMP,
                             deleted_at TIMESTAMP
);

CREATE TABLE fundraising_sponsor (
                                    id SERIAL PRIMARY KEY,
                                    fundraising_id UUID REFERENCES fundraising(id),
                                    user_id UUID REFERENCES user(id),
                                    status VARCHAR(50) NOT NULL,
                                    amount FLOAT NOT NULL,
                                    currency VARCHAR(10) NOT NULL,
                                    created_at TIMESTAMP,
                                    updated_at TIMESTAMP,
                                    deleted_at TIMESTAMP
);

CREATE TABLE country (
                         id SERIAL PRIMARY KEY,
                         name VARCHAR(255) NOT NULL,
                         code VARCHAR(50) NOT NULL,
                         lang VARCHAR(50) NOT NULL,
                         created_at TIMESTAMPTZ NOT NULL,
                         updated_at TIMESTAMPTZ NOT NULL,
                         deleted_at TIMESTAMPTZ
);
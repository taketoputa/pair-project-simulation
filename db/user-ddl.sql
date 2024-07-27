-- table1
create table products (
    id int auto_increment primary key,
    name varchar(100) not null,
    price decimal(10, 2) not null,
    stock int not null
);

-- table2
create table staff (
    id int auto_increment primary key,
    name varchar(100) not null,
    email varchar(100) not null unique,
    position varchar(50) not null
);

-- table3
create table sales (
    id int auto_increment primary key,
    product_id int not null,
    quantity int not null,
    sale_date datetime not null,
    foreign key (product_id) references products(id)
);

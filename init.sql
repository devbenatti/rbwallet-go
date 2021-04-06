CREATE TABLE operations (
  `id` 				INT(11)			NOT NULL AUTO_INCREMENT,
  `description` 			VARCHAR(100)				NOT NULL,
  `updated_at` TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
  `created_at` TIMESTAMP NOT NULL DEFAULT NOW(),
  PRIMARY KEY (`id`)
);


CREATE TABLE `accounts` (
  `id` varchar(36) NOT NULL,
  `document_identifier` varchar(14) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE transactions (
  `code` 					VARCHAR(36)			NOT NULL,
  `account_id`              VARCHAR(36)             NOT NULL,
  `operation_id`            INT(11)             NOT NULL,
  `total`  					DECIMAL(15,2)       NOT NULL,
  `updated_at` TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
  `created_at` TIMESTAMP NOT NULL DEFAULT NOW(),
  PRIMARY KEY (`code`),
  FOREIGN KEY (`account_id`) REFERENCES accounts (`id`),
  FOREIGN KEY (`operation_id`) REFERENCES operations (`id`)
);

INSERT INTO onepuchman.operations
(id, description)
VALUES(1, 'COMPRA A VISTA');
INSERT INTO onepuchman.operations
(id, description)
VALUES(2, 'COMPRA PARCELADA');
INSERT INTO onepuchman.operations
(id, description)
VALUES(3, 'SAQUE');
INSERT INTO onepuchman.operations
(id, description)
VALUES(4, 'PAGAMENTO');


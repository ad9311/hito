INSERT INTO users (
	"name",
	username,
	password,
	admin,
	created_at,
	updated_at
)
VALUES (
	'Default User',
	'default',
	'$2a$10$DdD4UVzhZRmLl3OXC8ZkvufWyWVk8qVkXR9AppKxbEBc.em.4/cY.',
	'false',
	'2021-12-04 19:39:00',
	'2021-12-04 19:39:00'
);

INSERT INTO landmarks (
	"name",
	native_name,
	"type",
	description,
	continent,
	country,
	state_city,
	start_year,
	end_year,
	area,
	height,
	wiki_url,
	img_url,
	created_at,
	updated_at
	)
VALUES (
	'San Felipe de Barajas Castle',
	'Castillo de San Felipe de Barajas',
	'Fortification',
	'The Castillo San Felipe de Barajas is a fortress in the city of Cartagena, Colombia. The castle is located on the Hill of San Lázaro in a strategic location, dominating approaches to the city by land or sea.',
	'South America',
	'Colombia',
	'Cartagena',
	1536,
	1767,
	142.28,
	42,
	'https://en.wikipedia.org/wiki/Castillo_San_Felipe_de_Barajas',
	'https://upload.wikimedia.org/wikipedia/commons/thumb/b/bc/62_-_Carthag%C3%A8ne_-_D%C3%A9cembre_2008.jpg/640px-62_-_Carthag%C3%A8ne_-_D%C3%A9cembre_2008.jpg',
	'2021-12-01 1:30',
	'2021-12-01 1:30'
);

package repository

var (
	queryExistsId string = `SELECT estado FROM lealtest.users
								WHERE EXISTS 
									(SELECT nombre FROM lealtest.users WHERE id=?);`
	userQuery            string = `SELECT * FROM lealtest.users AS u WHERE u.id = ? AND u.nombre = ?;`
	branchActivatedQuery string = `SELECT estado_campaña FROM lealtest.sucursales AS s 
							WHERE s.estado_campaña= 1 AND  s.id =?;`
)

package repository

var (
	queryExistsUserId string = `SELECT estado FROM lealtest.users
								WHERE EXISTS 
									(SELECT nombre FROM lealtest.users WHERE id=?);`
	nameQuery            string = `SELECT nombre FROM lealtest.users WHERE id=?;`
	userQuery            string = `SELECT * FROM lealtest.users AS u WHERE u.id = ? AND u.nombre = ?;`
	branchActivatedQuery string = `SELECT estado_campaña FROM lealtest.sucursales 	AS s 
							WHERE s.estado_campaña= 1 AND  s.id =?;`
	campaignNameQuery string = `SELECT c.nombre FROM lealtest.sucursales AS s 
	INNER JOIN lealtest.campañas AS c ON s.campaña_id = c.id
    WHERE s.id = ?;`
	updateUserCashbackPoints string = `UPDATE lealtest.users AS u
				SET u.puntos = ?, u.cashback = ?, u.updatedAt = NOW()
				WHERE u.id =?  AND u.nombre = ?;`
)

package repository

var (
	createQuery string = `INSERT INTO lealtest.campañas (nombre, fecha_activacion, fecha_inactivacion, estado, sucursal, descripcion, createdAt, UpdatedAt) 
								values (?,?,?,?,?,?,?,?);`
	getByBranchQuery   string = `SELECT * FROM lealtest.campañas AS c WHERE c.sucursal = ? AND c.estado = ?;`
	getByCommerceQuery string = ` SELECT c.id, c.nombre, c.fecha_activacion, c.fecha_inactivacion, c.estado, c.sucursal,
											c.descripcion, c.createdAt, c.updatedAt
									FROM lealtest.campañas AS c
										INNER JOIN lealtest.sucursales AS s
											ON c.sucursal = s.id
										WHERE s.comercio_id = ? AND c.estado=?;`
	activateCampaignQuery string = `UPDATE lealtest.campañas AS c 
									SET c.fecha_activacion= NOW(), c.estado=1, c.updatedAt=NOW() 
									WHERE c.nombre = ? AND c.sucursal=?;`
	inactivateCampaignQuery string = `UPDATE lealtest.campañas AS c 
									SET c.fecha_inactivacion= NOW(), c.estado=0, c.updatedAt=NOW() 
									WHERE c.nombre = ? AND c.sucursal=?;`
	campaignActiveInBranch string = `SELECT estado_campaña FROM lealtest.sucursales AS s WHERE s.id=? AND s.estado_campaña = 1; `
	branchCampaignActivate string = `UPDATE lealtest.sucursales AS s 
									  SET s.estado_campaña= 1, s.updatedAt=NOW() WHERE s.id = ?;`
	updateCommerceIdInBranch string = `UPDATE lealtest.sucursales AS s
										INNER JOIN  lealtest.campañas AS c ON s.id = ?
										SET s.campaña_id = c.id
										WHERE c.estado = 1 AND c.sucursal = ?;`
	branchCampaignInactivate string = `UPDATE lealtest.sucursales AS s
											SET s.estado_campaña = 0, s.updatedAt=NOW()
											WHERE s.id = ?;`
	campaignStatusQuery string = `SELECT campaña_id, c.nombre, fecha_activacion, fecha_inactivacion FROM lealtest.sucursales AS s 
									INNER JOIN lealtest.campañas AS c ON s.campaña_id = c.id
									WHERE s.estado_campaña= 1 AND  s.id =?;`
)

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
									SET c.fecha_activacion= NOW(), c.estado=1 
									WHERE c.nombre = ? AND c.sucursal=?;`
	inactivateCampaignQuery string = `UPDATE lealtest.campañas AS c 
									SET c.fecha_inactivacion= NOW(), c.estado=0 
									WHERE c.nombre = ? AND c.sucursal=?;`
)

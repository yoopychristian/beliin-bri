package database

import (
	"gorm.io/gorm"
)

const DATASIZE = 25

type UserRegistration struct {
	IDUser    string `gorm:"column:id_user"`
	Nama      string `gorm:"column:nama"`
	Username  string `gorm:"column:username"`
	Password  string `gorm:"column:password;type:varchar;size:255"`
	Email     string `gorm:"column:email"`
	NoPonsel  string `gorm:"column:no_ponsel"`
	NoKTP     string `gorm:"column:no_ktp"`
	IsDeleted int    `gorm:"column:is_deleted"`
}

// func (p UserRegistration) TableName() string {
// 	return "user-registration"
// }

func (p *UserRegistration) GetByUsername(db *gorm.DB, username string) error {
	return db.Table("user_registration").Where("username=?", username).Last(&p).Error
}

// func (p *Auth) List(db *gorm.DB, username, contact, email string, page int) ([]shared.AuthDetail, error) {
// 	errHandle := func(err error) ([]shared.AuthDetail, error) {
// 		return nil, err
// 	}

// 	data := []shared.AuthDetail{}
// 	offset := 0
// 	if page > 1 {
// 		offset = DATASIZE * (page - 1)
// 	}

// 	format := "select a.id,a.username,a.full_name,a.email,a.phone,a.created_at,a.login_at,"
// 	format += "a.login_ip,a.last_seen,a.user_agent,b.group_name,a.access_group "
// 	format += "from admin as a inner join admin_group as b on a.access_group =b.id  "
// 	format += "where a.deleted_at isnull [filter] order by a.username limit %d offset %d"

// 	var rows *sql.Rows

// 	//non filter
// 	if username != "" {
// 		sql := fmt.Sprintf(format, DATASIZE, offset)
// 		sql = strings.ReplaceAll(sql, "[filter]", "and username like ?")

// 		var err error
// 		rows, err = db.Raw(sql, "%"+username+"%").Rows()
// 		if err != nil {
// 			return errHandle(err)
// 		}
// 	} else if email != "" {
// 		sql := fmt.Sprintf(format, DATASIZE, offset)
// 		sql = strings.ReplaceAll(sql, "[filter]", "and email like ?")

// 		var err error
// 		rows, err = db.Raw(sql, "%"+email+"%").Rows()
// 		if err != nil {
// 			return errHandle(err)
// 		}
// 	} else if contact != "" {
// 		sql := fmt.Sprintf(format, DATASIZE, offset)
// 		sql = strings.ReplaceAll(sql, "[filter]", "and phone like ?")

// 		var err error
// 		rows, err = db.Raw(sql, "%"+contact+"%").Rows()
// 		if err != nil {
// 			return errHandle(err)
// 		}
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		p := shared.AdminDetail{}
// 		rows.Scan(
// 			&p.ID,
// 			&p.Username,
// 			&p.FullName,
// 			&p.Email,
// 			&p.Phone,
// 			&p.CreatedAt,
// 			&p.LoginAt,
// 			&p.IP,
// 			&p.LastSeen,
// 			&p.UserAgent,
// 			&p.GroupName,
// 			&p.GroupID,
// 		)
// 		p.Created = p.CreatedAt.Format("2006-01-02 15:04:05")
// 		p.Last = p.LastSeen.Format("2006-01-02 15:04:05")
// 		p.Login = p.LoginAt.Format("2006-01-02 15:04:05")
// 		data = append(data, p)
// 	}

// 	return data, nil
// }

// func (p Admin) Count(db *gorm.DB) int64 {
// 	var count int64
// 	db.Model(&Admin{}).Where("deleted_at isnull").Count(&count)
// 	return count
// }

// func (p Admin) TotalPage(totalRows int) int {
// 	return (totalRows / DATASIZE) + 1
// }

package server

import "net/http"

type Authorization struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Desc     string `json:"description"`
	CreateAt string `json:"create_at"`
}

func (s *Server) handleAuthorizationList(
	w http.ResponseWriter,
	r *http.Request,
) {

	user, err := s.auth.GetUser(r)
	if err != nil {
		resp(w, 10002, err)
		return
	}

	rows, err := s.db.QueryContext(
		dbCtx(),
		"SELECT `user_auth`.`id`,`name`,`description`,`app`.`create_at` "+
			"FROM `user_auth` LEFT JOIN `app` "+
			"ON `user_auth`.`app_id`=`app`.`app_id` "+
			"WHERE `username`=? ORDER BY `user_auth`.`create_at` DESC "+
			"LIMIT 100",
		user.Username,
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	list := []*Authorization{}

	for rows.Next() {
		auth := &Authorization{}
		if err := rows.Scan(
			&auth.ID,
			&auth.Name,
			&auth.Desc,
			&auth.CreateAt,
		); err != nil {
			panic(err)
		}
		list = append(list, auth)
	}

	resp(w, 0, map[string]interface{}{
		"list": list,
	})
}

type AuthorizationDeleteParam struct {
	ID string `json:"id"`
}

func (s *Server) handleAuthorizationDelete(
	w http.ResponseWriter,
	r *http.Request,
) {

	user, err := s.auth.GetUser(r)
	if err != nil {
		resp(w, 10002, err)
		return
	}
	in := &AuthorizationDeleteParam{}
	if err := bind(r, in); err != nil {
		resp(w, 10001, err)
		return
	}

	if _, err := s.db.ExecContext(
		dbCtx(),
		"DELETE FROM `user_auth` WHERE `id`=? AND `user_id`=? LIMIT 1",
		in.ID,
		user.ID,
	); err != nil {
		panic(err)
	}
	resp(w, 0, nil)
}

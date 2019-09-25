package server

import "net/http"

type AppCreateParam struct {
	Name            string `json:"name"`
	Desc            string `json:"description"`
	OIDCRedirectURI string `json:"oidc_redirect_uri"`
}

func (s *Server) handleAppCreate(w http.ResponseWriter, r *http.Request) {
	user, err := s.auth.GetUser(r)
	if err != nil {
		resp(w, 10002, err)
		return
	}
	in := &AppCreateParam{}
	if err := bind(r, in); err != nil {
		resp(w, 10001, err)
		return
	}

	if in.Name == "" {
		resp(w, 10001, errAPPNameInvalid)
		return
	}

	if _, err := s.db.ExecContext(
		dbCtx(),
		"INSERT INTO `app` "+
			"(`app_id`, `secret`, `name`, `description`, "+
			"`oidc_redirect_uri`, `user_id`) "+
			"VALUES(?,?,?,?,?,?)",
		RandToken(),
		RandToken(),
		in.Name,
		in.Desc,
		in.OIDCRedirectURI,
		user.ID,
	); err != nil {
		panic(err)
	}
	resp(w, 0, nil)
}

type AppUpdateParam struct {
	AppID           string `json:"app_id"`
	Name            string `json:"name"`
	Desc            string `json:"description"`
	OIDCRedirectURI string `json:"oidc_redirect_uri"`
}

func (s *Server) handleAppUpdate(w http.ResponseWriter, r *http.Request) {
	user, err := s.auth.GetUser(r)
	if err != nil {
		resp(w, 10002, err)
		return
	}
	in := &AppUpdateParam{}
	if err := bind(r, in); err != nil {
		resp(w, 10001, err)
		return
	}

	if in.Name == "" {
		resp(w, 10001, errAPPNameInvalid)
		return
	}

	if _, err := s.db.ExecContext(
		dbCtx(),
		"UPDATE `app` SET "+
			"`name`=?,`description`=?,`oidc_redirect_uri`=?,`update_at`=NOW() "+
			"WHERE `app_id`=? AND `user_id`=? LIMIT 1",
		in.Name,
		in.Desc,
		in.OIDCRedirectURI,
		in.AppID,
		user.ID,
	); err != nil {
		panic(err)
	}
	resp(w, 0, nil)
}

type AppDeleteParam struct {
	AppID string `json:"app_id"`
}

func (s *Server) handleAppDelete(w http.ResponseWriter, r *http.Request) {
	user, err := s.auth.GetUser(r)
	if err != nil {
		resp(w, 10002, err)
		return
	}
	in := &AppDeleteParam{}
	if err := bind(r, in); err != nil {
		resp(w, 10001, err)
		return
	}

	if _, err := s.db.ExecContext(
		dbCtx(),
		"DELETE FROM `app` WHERE `app_id`=? AND `user_id`=? LIMIT 1",
		in.AppID,
		user.ID,
	); err != nil {
		panic(err)
	}
	resp(w, 0, nil)
}

type App struct {
	AppID           string `json:"app_id"`
	Secret          string `json:"secret"`
	Name            string `json:"name"`
	Desc            string `json:"description"`
	OIDCRedirectURI string `json:"oidc_redirect_uri"`
	UpdateAt        string `json:"update_at"`
}

func (s *Server) handleAppList(w http.ResponseWriter, r *http.Request) {
	user, err := s.auth.GetUser(r)
	if err != nil {
		resp(w, 10002, err)
		return
	}

	rows, err := s.db.QueryContext(
		dbCtx(),
		"SELECT `app_id`, `secret`, `name`, `description`, "+
			"`oidc_redirect_uri`, `update_at` "+
			"FROM `app` WHERE `user_id` = ? LIMIT 100",
		user.ID,
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	list := []*App{}

	for rows.Next() {
		app := &App{}
		if err := rows.Scan(
			&app.AppID,
			&app.Secret,
			&app.Name,
			&app.Desc,
			&app.OIDCRedirectURI,
			&app.UpdateAt,
		); err != nil {
			panic(err)
		}
		list = append(list, app)
	}

	resp(w, 0, map[string]interface{}{
		"list": list,
	})
}

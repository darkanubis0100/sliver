package privilege

/*
	Sliver Implant Framework
	Copyright (C) 2021  Bishop Fox

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"context"

	"github.com/bishopfox/sliver/client/console"
	"github.com/bishopfox/sliver/protobuf/sliverpb"
	"github.com/desertbit/grumble"
)

func ImpersonateCmd(ctx *grumble.Context, con *console.SliverConsoleClient) {
	session := con.ActiveSession.GetInteractive()
	if session == nil {
		return
	}
	username := ctx.Args.String("username")
	impResp, err := con.Rpc.Impersonate(context.Background(), &sliverpb.ImpersonateReq{
		Request:  con.ActiveSession.Request(ctx),
		Username: username,
	})

	if err != nil {
		con.PrintErrorf("%s", err)
		return
	}
	if impResp.GetResponse().GetErr() != "" {
		con.PrintErrorf("%s\n", impResp.GetResponse().GetErr())
		return
	}
	con.PrintInfof("Successfully impersonated %s\n", username)
}

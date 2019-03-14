package apps

import (
	"github.com/TeaWeb/code/teaconfigs/agents"
	"github.com/TeaWeb/code/teaconfigs/notices"
	"github.com/TeaWeb/code/teaweb/actions/default/agents/agentutils"
	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
)

type ItemDetailAction actions.Action

// 监控项详情
func (this *ItemDetailAction) Run(params struct {
	AgentId string
	AppId   string
	ItemId  string
}) {
	this.Data["agent"] = maps.Map{
		"id": params.AgentId,
	}

	app := agentutils.InitAppData(this, params.AgentId, params.AppId, "monitor")
	item := app.FindItem(params.ItemId)

	if item == nil {
		this.Fail("找不到要查看的Item")
	}

	// recover success
	if item.RecoverSuccesses <= 0 {
		item.RecoverSuccesses = 1
	}

	this.Data["item"] = item

	this.Data["sourceOptions"] = nil
	this.Data["sourcePresentation"] = nil
	this.Data["source"] = nil
	source := item.Source()
	if source != nil {
		summary := agents.FindDataSource(source.Code())
		summary["variables"] = source.Variables()
		this.Data["sourceOptions"] = maps.Map{
			"summary":    summary,
			"options":    source,
			"dataFormat": agents.FindSourceDataFormat(source.DataFormatCode()),
		}

		this.Data["source"] = source

		p := source.Presentation()
		if p != nil {
			p.CSS = "<style type=\"text/css\">\n" + p.CSS + "\n</style>\n"
			this.Data["sourcePresentation"] = p
		}
	}

	this.Data["noticeLevels"] = notices.AllNoticeLevels()

	this.Show()
}

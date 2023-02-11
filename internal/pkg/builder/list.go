package builder

import (
	agent "go-to-cloud/internal/agent_server"
	"go-to-cloud/internal/models"
	"go-to-cloud/internal/models/builder"
	"go-to-cloud/internal/repositories"
	"go-to-cloud/internal/utils"
)

func ListNodesOnK8s(orgs []uint, query *builder.Query) ([]builder.NodesOnK8s, error) {
	var orgId []uint
	if query == nil || len(query.Orgs) == 0 {
		//	默认取当前用户所属全体组织
		orgId = orgs
	} else {
		// 	计算查询条件中的所属组织与当前用户所属组织的交集
		orgId = utils.Intersect(orgs, query.Orgs)
	}

	patternName := ""
	pager := models.Pager{}
	if query != nil {
		patternName = query.Name
		pager = query.Pager
	}
	if merged, err := repositories.GetBuildNodesOnK8sByOrgId(orgId, patternName, &pager); err != nil {
		return nil, err
	} else {
		rlt := make([]builder.NodesOnK8s, len(merged))
		for i, m := range merged {
			orgLites := make([]builder.OrgLite, len(m.Org))
			for i, lite := range m.Org {
				orgLites[i] = builder.OrgLite{
					OrgId:   lite.OrgId,
					OrgName: lite.OrgName,
				}
			}
			rlt[i] = builder.NodesOnK8s{
				Id:           m.ID,
				Name:         m.Name,
				OrgLites:     orgLites,
				Remark:       m.Remark,
				AgentVersion: m.AgentVersion,
				Workspace:    m.K8sWorkerSpace,
				MaxWorkers:   m.MaxWorkers,
				KubeConfig:   "***Hidden***", // func() string { return *m.DecryptKubeConfig() }(),
				CurrentWorkers: func() int {
					return (agent.Runner).GetNodeCount(int64(m.ID))
				}(),
			}
		}
		return rlt, err
	}
}

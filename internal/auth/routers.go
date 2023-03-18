package auth

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/conf"
	"go-to-cloud/internal/controllers/auth"
	"go-to-cloud/internal/controllers/configure/artifact"
	"go-to-cloud/internal/controllers/configure/buildEnv"
	"go-to-cloud/internal/controllers/configure/builder"
	"go-to-cloud/internal/controllers/configure/deploy/k8s"
	"go-to-cloud/internal/controllers/configure/scm"
	"go-to-cloud/internal/controllers/monitor"
	"go-to-cloud/internal/controllers/projects"
	"go-to-cloud/internal/controllers/users"
)

type RestfulMethod string

const (
	PUT    RestfulMethod = "PUT"
	GET    RestfulMethod = "GET"
	DELETE RestfulMethod = "DELETE"
	POST   RestfulMethod = "POST"
)

type RouterMap struct {
	Url     string
	Methods []RestfulMethod
	Func    func(ctx *gin.Context)
	Kinds   []conf.Kind
}

var RouterMaps []RouterMap

func init() {
	RouterMaps = make([]RouterMap, 0)

	// webapi
	RouterMaps = append(RouterMaps, RouterMap{"/api/user/auths", []RestfulMethod{GET}, auth.GetAuthCodes, []conf.Kind{conf.Guest}})                    // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/user/kinds", []RestfulMethod{GET}, users.AllKinds, []conf.Kind{conf.Guest}})                       // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/user", []RestfulMethod{PUT, POST}, users.UpsertUser, []conf.Kind{conf.Guest}})                     // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/user/info", []RestfulMethod{GET}, users.Info, []conf.Kind{conf.Guest}})                            // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/user/org/list", []RestfulMethod{GET}, users.OrgList, []conf.Kind{conf.Guest}})                     // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/user/org", []RestfulMethod{PUT, POST}, users.UpsertOrg, []conf.Kind{conf.Guest}})                  // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/user/org/:orgId", []RestfulMethod{DELETE}, users.DeleteOrg, []conf.Kind{conf.Guest}})              // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/user/list", []RestfulMethod{GET}, users.List, []conf.Kind{conf.Guest}})                            // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/user/joined/:orgId", []RestfulMethod{GET}, users.Joined, []conf.Kind{conf.Guest}})                 // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/user/join/:orgId", []RestfulMethod{PUT}, users.Join, []conf.Kind{conf.Guest}})                     // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/user/:userId/orgs/joined", []RestfulMethod{GET}, users.Belonged, []conf.Kind{conf.Guest}})         // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/user/:userId/join", []RestfulMethod{PUT}, users.Belong, []conf.Kind{conf.Guest}})                  // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/user/:userId", []RestfulMethod{DELETE}, users.DeleteUser, []conf.Kind{conf.Guest}})                // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/user/:userId/password/reset", []RestfulMethod{PUT}, users.ResetPassword, []conf.Kind{conf.Guest}}) // 所有角色默认具备Guest的所有权限

	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/build/env", []RestfulMethod{GET}, buildEnv.BuildEnv, []conf.Kind{conf.Guest}}) // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/build/cmd", []RestfulMethod{GET}, buildEnv.BuildCmd, []conf.Kind{conf.Guest}}) // 所有角色默认具备Guest的所有权限

	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/coderepo", []RestfulMethod{GET}, scm.QueryCodeRepos, []conf.Kind{conf.Guest}})        // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/coderepo/bind", []RestfulMethod{POST}, scm.BindCodeRepo, []conf.Kind{conf.Guest}})    // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/coderepo", []RestfulMethod{PUT}, scm.UpdateCodeRepo, []conf.Kind{conf.Guest}})        // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/coderepo/:id", []RestfulMethod{DELETE}, scm.RemoveCodeRepo, []conf.Kind{conf.Guest}}) // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/coderepo/testing", []RestfulMethod{POST}, scm.Testing, []conf.Kind{conf.Guest}})      // 所有角色默认具备Guest的所有权限

	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/artifact/testing", []RestfulMethod{POST}, artifact.Testing, []conf.Kind{conf.Guest}})                             // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/artifact/bind", []RestfulMethod{POST}, artifact.BindArtifactRepo, []conf.Kind{conf.Guest}})                       // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/artifact", []RestfulMethod{PUT}, artifact.UpdateArtifactRepo, []conf.Kind{conf.Guest}})                           // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/artifact", []RestfulMethod{GET}, artifact.QueryArtifactRepo, []conf.Kind{conf.Guest}})                            // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/artifact/:id", []RestfulMethod{DELETE}, artifact.RemoveArtifactRepo, []conf.Kind{conf.Guest}})                    // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/artifact/:id", []RestfulMethod{GET}, artifact.QueryArtifactItems, []conf.Kind{conf.Guest}})                       // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/artifact/image/:imageId", []RestfulMethod{DELETE}, artifact.DeleteImage, []conf.Kind{conf.Guest}})                // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/artifact/images/hashId/:hashId", []RestfulMethod{DELETE}, artifact.DeleteImageByHashId, []conf.Kind{conf.Guest}}) // 所有角色默认具备Guest的所有权限

	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/deploy/k8s/testing", []RestfulMethod{POST}, k8s.Testing, []conf.Kind{conf.Guest}})     // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/deploy/k8s/bind", []RestfulMethod{POST}, k8s.BindK8sRepo, []conf.Kind{conf.Guest}})    // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/deploy/k8s", []RestfulMethod{PUT}, k8s.UpdateK8sRepo, []conf.Kind{conf.Guest}})        // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/deploy/k8s", []RestfulMethod{GET}, k8s.QueryK8sRepos, []conf.Kind{conf.Guest}})        // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/deploy/k8s/:id", []RestfulMethod{DELETE}, k8s.RemoveK8sRepo, []conf.Kind{conf.Guest}}) // 所有角色默认具备Guest的所有权限

	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/builder/install/k8s", []RestfulMethod{POST}, builder.K8sInstall, []conf.Kind{conf.Guest}})                      // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/builder/nodes/k8s", []RestfulMethod{GET}, builder.QueryNodesOnK8s, []conf.Kind{conf.Guest}})                    // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/builder/nodes/k8s/available", []RestfulMethod{GET}, builder.QueryAvailableNodesOnK8s, []conf.Kind{conf.Guest}}) // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/builder/node/:id", []RestfulMethod{DELETE}, builder.Uninstall, []conf.Kind{conf.Guest}})                        // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/configure/builder/node", []RestfulMethod{PUT}, builder.UpdateBuilderNode, []conf.Kind{conf.Guest}})                       // 所有角色默认具备Guest的所有权限

	RouterMaps = append(RouterMaps, RouterMap{"/api/projects", []RestfulMethod{POST}, projects.Create, []conf.Kind{conf.Guest}})                                                            // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId", []RestfulMethod{DELETE}, projects.DeleteProject, []conf.Kind{conf.Guest}})                                        // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/list", []RestfulMethod{GET}, projects.List, []conf.Kind{conf.Guest}})                                                          // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/coderepo", []RestfulMethod{GET}, projects.CodeRepo, []conf.Kind{conf.Guest}})                                                  // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects", []RestfulMethod{PUT}, projects.UpdateProject, []conf.Kind{conf.Guest}})                                                      // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/import", []RestfulMethod{POST}, projects.ImportSourceCode, []conf.Kind{conf.Guest}})                                // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/imported", []RestfulMethod{GET}, projects.ListImportedSourceCode, []conf.Kind{conf.Guest}})                         // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/sourcecode/:id", []RestfulMethod{DELETE}, projects.DeleteSourceCode, []conf.Kind{conf.Guest}})                      // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/src/:sourceCodeId", []RestfulMethod{GET}, projects.ListBranches, []conf.Kind{conf.Guest}})                          // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/pipeline", []RestfulMethod{POST}, projects.NewBuildPlan, []conf.Kind{conf.Guest}})                                  // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/pipeline", []RestfulMethod{GET}, projects.QueryBuildPlan, []conf.Kind{conf.Guest}})                                 // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/pipeline/:pipelineId/history", []RestfulMethod{GET}, projects.QueryBuildPlanHistory, []conf.Kind{conf.Guest}})      // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/pipeline/state", []RestfulMethod{GET}, projects.QueryBuildPlanState, []conf.Kind{conf.Guest}})                      // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/pipeline/:id", []RestfulMethod{DELETE}, projects.DeleteBuildPlan, []conf.Kind{conf.Guest}})                         // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/pipeline/:id/build", []RestfulMethod{POST}, projects.StartBuildPlan, []conf.Kind{conf.Guest}})                      // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/deploy/:id", []RestfulMethod{PUT}, projects.Deploying, []conf.Kind{conf.Guest}})                                    // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/deploy/:id/rollback/:historyId", []RestfulMethod{PUT}, projects.Rollback, []conf.Kind{conf.Guest}})                 // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/deploy/apps", []RestfulMethod{GET}, projects.QueryDeployments, []conf.Kind{conf.Guest}})                            // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/deploy/app/:deploymentId/history", []RestfulMethod{GET}, projects.QueryDeploymentHistory, []conf.Kind{conf.Guest}}) // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/deploy/app", []RestfulMethod{POST}, projects.CreateDeployment, []conf.Kind{conf.Guest}})                            // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/deploy/:id", []RestfulMethod{DELETE}, projects.DeleteDeployment, []conf.Kind{conf.Guest}})                          // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/deploy/:k8sRepoId/namespaces", []RestfulMethod{GET}, projects.QueryNamespaces, []conf.Kind{conf.Guest}})            // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/deploy/env", []RestfulMethod{GET}, projects.QueryDeploymentEnv, []conf.Kind{conf.Guest}})                           // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/artifacts/:querystring", []RestfulMethod{GET}, projects.QueryArtifacts, []conf.Kind{conf.Guest}})                   // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/artifact/:artifactId/tags", []RestfulMethod{GET}, projects.QueryArtifactTags, []conf.Kind{conf.Guest}})             // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/artifact/:artifactId", []RestfulMethod{GET}, projects.QueryArtifactItemsByProjectId, []conf.Kind{conf.Guest}})      // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/projects/:projectId/artifact", []RestfulMethod{GET}, projects.QueryArtifactsByProjectId, []conf.Kind{conf.Guest}})                      // 所有角色默认具备Guest的所有权限

	RouterMaps = append(RouterMaps, RouterMap{"/api/monitor/:k8s/apps/query", []RestfulMethod{GET}, monitor.Query, []conf.Kind{conf.Guest}})                              // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/monitor/:k8s/apps/restart", []RestfulMethod{PUT}, monitor.Restart, []conf.Kind{conf.Guest}})                          // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/monitor/:k8s/apps/delete", []RestfulMethod{PUT}, monitor.DeletePod, []conf.Kind{conf.Guest}})                         // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/monitor/:k8s/apps/scale", []RestfulMethod{PUT}, monitor.Scale, []conf.Kind{conf.Guest}})                              // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/monitor/:k8s/pods/:deploymentId", []RestfulMethod{GET}, monitor.QueryPods, []conf.Kind{conf.Guest}})                  // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/api/monitor/:k8s/apps/delete/:deploymentId", []RestfulMethod{DELETE}, monitor.DeleteDeployment, []conf.Kind{conf.Guest}}) // 所有角色默认具备Guest的所有权限

	// websocket
	RouterMaps = append(RouterMaps, RouterMap{"/ws/monitor/:k8s/pod/:deploymentId/:podName/log", []RestfulMethod{GET}, monitor.DisplayLog, []conf.Kind{conf.Guest}})    // 所有角色默认具备Guest的所有权限
	RouterMaps = append(RouterMaps, RouterMap{"/ws/monitor/:k8s/pod/:deploymentId/:podName/shell", []RestfulMethod{GET}, monitor.Interactive, []conf.Kind{conf.Guest}}) // 所有角色默认具备Guest的所有权限
}

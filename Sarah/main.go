// MIT-License https://github.com/higashi000/procon30-kyogi
package main

import (
  "github.com/gin-gonic/gin"
  "io/ioutil"
  "os"
  "log"
  "fmt"
  "encoding/json"
  "time"
)

type FieldData struct {
	Width             int     `json:"width"`
	Height            int     `json:"height"`
	Points            [][]int `json:"points"`
	StartedAtUnixTime int     `json:"startedAtUnixTime"`
	Turn              int     `json:"turn"`
	Tiled             [][]int `json:"tiled"`
	Teams             []struct {
		TeamID int `json:"teamID"`
		Agents []struct {
			AgentID int `json:"agentID"`
			X       int `json:"x"`
			Y       int `json:"y"`
		} `json:"agents"`
		TilePoint int `json:"tilePoint"`
		AreaPoint int `json:"areaPoint"`
	} `json:"teams"`
	Actions []interface{} `json:"actions"`
}

type Test struct {
  Num int `json:num`
}

// エージェントの行動を受け取る構造体 {{{
type  Action struct {
  AgentID int `json:"agentID"`
  Type string `json:"type"`
  Dx int `json:"dx"`
  Dy int `json:"dy"`
}

type Actions struct {
  AgentActions []Action `json:"actions"`
}
// }}}

// jsonファイルの読み込み
func getFieldData() FieldData {
  // ファイルから読み取り
  bytes, err := ioutil.ReadFile("A.json")

  if err != nil {
    log.Fatal(err)
  }

  var field FieldData

    // fieldのstructに格納
  if err := json.Unmarshal(bytes, &field); err != nil {
    log.Fatal(err)
  }

  return field
}

func main() {
  turn := 1
  r := gin.Default()
  field := getFieldData()
  field = reLoadFieldData(r, field)
  sendFieldData(r, field)
  rsvActionData(r, &field, &turn)
  r.Run()
}

// フィールドデータの返却
func sendFieldData(r *gin.Engine, field FieldData) {
  r.GET("/matches/:id", func(c *gin.Context) {
      c.JSON(200, field)
      })
}

func reLoadFieldData(r *gin.Engine, field FieldData) (FieldData) {
  r.POST("/reload/:id", func(c *gin.Context) {
      os.Exit(1)
      })

  return field
}

func rsvActionData(r *gin.Engine, field *FieldData, turn *int) {
  var actions Actions
    r.POST("/matches/:id/action", func(c *gin.Context) {
        c.BindJSON(&actions)
        for i := 0; i < len(field.Teams[0].Agents); i++ {
        field.Teams[0].Agents[i].X = field.Teams[0].Agents[i].X
        field.Teams[0].Agents[i].Y = field.Teams[0].Agents[i].Y
        field.Teams[1].Agents[i].X = field.Teams[1].Agents[i].X
        field.Teams[1].Agents[i].Y = field.Teams[1].Agents[i].Y
        }
        if field.Teams[0].Agents[0].AgentID == actions.AgentActions[0].AgentID {
        field.Turn += 1
        updateFieldData(field, actions, 1)
        } else {
        field.Turn += 1
        updateFieldData(field, actions, 2)
        }
        fmt.Println(field.Teams[1].Agents)
        *field = getFieldData()
        })
}

func checkDuplicate(whichTeam bool, tmpPos [][]int, agentNum int, field FieldData) {
  if whichTeam {
    for i := 0; i < agentNum; i++ {
      for j := 0; j < len(tmpPos); j++ {
        if i != j && tmpPos[i][0] == tmpPos[j][0] && tmpPos[i][1] == tmpPos[j][1] {
          tmpPos[i][0] = field.Teams[0].Agents[i].X - 1
            tmpPos[i][1] = field.Teams[0].Agents[i].Y - 1

            if j < agentNum {
              tmpPos[j][0] = field.Teams[0].Agents[j].X - 1
                tmpPos[j][1] = field.Teams[0].Agents[j].Y - 1
            }
        }
      }
    }
  } else {
    for i := 0; i < agentNum; i++ {
      for j := 0; j < len(tmpPos); j++ {
        if i != j && tmpPos[i][0] == tmpPos[j][0] && tmpPos[i][1] == tmpPos[j][1] {
          tmpPos[i][0] = field.Teams[1].Agents[i].X - 1
            tmpPos[i][1] = field.Teams[1].Agents[i].Y - 1

            if j < agentNum {
              tmpPos[j][0] = field.Teams[0].Agents[j].X - 1
                tmpPos[j][1] = field.Teams[0].Agents[j].Y - 1
            }
        }
      }
    }
  }
}

func updateFieldData(field *FieldData, action Actions, whichTeam int) {
agentNum := len(field.Teams[0].Agents)
            tmpPos := make([][]int, agentNum * 2)

            for i := 0; i < agentNum * 2; i++ {
              tmpPos[i] = []int{0, 0}
            }

          if whichTeam == field.Teams[0].TeamID {
            for i := 0; i < agentNum; i++ {
              tmpPos[i][0] = field.Teams[0].Agents[i].X + action.AgentActions[i].Dx - 1
                tmpPos[i][1] = field.Teams[0].Agents[i].Y + action.AgentActions[i].Dy - 1

                if tmpPos[i][0] < 0 || field.Width <= tmpPos[i][0] {
                  tmpPos[i][0] = field.Teams[0].Agents[i].X - 1
                    tmpPos[i][1] = field.Teams[0].Agents[i].Y - 1
                } else if tmpPos[i][1] < 0 || field.Height <= tmpPos[i][1] {
                  tmpPos[i][0] = field.Teams[0].Agents[i].X - 1
                    tmpPos[i][1] = field.Teams[0].Agents[i].Y - 1
                }

              if field.Tiled[tmpPos[i][1]][tmpPos[i][0]] == field.Teams[1].TeamID {
tmp := []int{field.Teams[0].Agents[i].X - 1, field.Teams[0].Agents[i].Y - 1}
     tmpPos = append(tmpPos, tmp)
              }
            }

            checkDuplicate(true, tmpPos, agentNum, *field)
              checkDuplicate(true, tmpPos, agentNum, *field)

              for i := 0; i < agentNum; i++ {
                if field.Tiled[tmpPos[i][1]][tmpPos[i][0]] == field.Teams[1].TeamID {
                  field.Tiled[tmpPos[i][1]][tmpPos[i][0]] = 0
                } else {
                  field.Teams[0].Agents[i].X = tmpPos[i][0] + 1
                    field.Teams[0].Agents[i].Y = tmpPos[i][1] + 1
                    field.Tiled[tmpPos[i][1]][tmpPos[i][0]] = field.Teams[0].TeamID
                }
              }
          } else {
            for i := 0; i < agentNum; i++ {
              tmpPos[i][0] = field.Teams[1].Agents[i].X + action.AgentActions[i].Dx - 1
                tmpPos[i][1] = field.Teams[1].Agents[i].Y + action.AgentActions[i].Dy - 1

                if tmpPos[i][0] < 0 || field.Width <= tmpPos[i][0] {
                  tmpPos[i][0] = field.Teams[1].Agents[i].X - 1
                    tmpPos[i][1] = field.Teams[1].Agents[i].Y - 1
                } else if tmpPos[i][1] < 0 || field.Height <= tmpPos[i][1] {
                  tmpPos[i][0] = field.Teams[1].Agents[i].X - 1
                    tmpPos[i][1] = field.Teams[1].Agents[i].Y - 1
                }

              if field.Tiled[tmpPos[i][1]][tmpPos[i][0]] == field.Teams[0].TeamID {
tmp := []int{field.Teams[1].Agents[i].X - 1, field.Teams[1].Agents[i].Y - 1}
     tmpPos = append(tmpPos, tmp)
              }
            }

            checkDuplicate(false, tmpPos, agentNum, *field)
              checkDuplicate(false, tmpPos, agentNum, *field)

              for i := 0; i < agentNum; i++ {
                if field.Tiled[tmpPos[i][1]][tmpPos[i][0]] == field.Teams[0].TeamID {
                  field.Tiled[tmpPos[i][1]][tmpPos[i][0]] = 0
                } else {
                  field.Teams[1].Agents[i].X = tmpPos[i][0] + 1
                    field.Teams[1].Agents[i].Y = tmpPos[i][1] + 1
                    field.Tiled[tmpPos[i][1]][tmpPos[i][0]] = field.Teams[1].TeamID
                }
              }
          }

tmpTeam1Tile := 0
                tmpTeam2Tile := 0
                for i := 0; i < field.Height; i++ {
                  for j := 0; j < field.Width; j++ {
                    if field.Tiled[i][j] == field.Teams[0].TeamID {
                      tmpTeam1Tile += field.Points[i][j]
                    } else if field.Tiled[i][j] == field.Teams[1].TeamID {
                      tmpTeam2Tile += field.Points[i][j]
                    }
                  }
                }

              field.Teams[0].TilePoint = tmpTeam1Tile;
              field.Teams[1].TilePoint = tmpTeam2Tile;

              field.StartedAtUnixTime = int(time.Now().Unix())
}

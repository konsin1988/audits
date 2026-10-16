package repositories

import (
    "database/sql"
    "time"
    "context"
    "konsin1988/go-server/models"
    _ "log"
)

type CkkCheckReestrRepo struct {
    db *sql.DB
}

func NewCkkCheckReestrRepo(db *sql.DB) *CkkCheckReestrRepo{
    return &CkkCheckReestrRepo{db: db}
}

func (r *CkkCheckReestrRepo) GetTableData(ctx context.Context) (models.TableData, error) {
    const query = `
        select 
            id,
            name,
            headCompNm,
            chArea,
            st_date,
            end_date,
            Responsible, 
            area_check
        from audit_ckk.ckkPlanGrafik
        `
    rows, err := r.db.QueryContext(ctx, query)
    if err!= nil {
        return models.TableData{}, err
    }
    defer rows.Close()

    data := models.TableData{
        Title:  "Реестр проверок",
        Headers: []string{"Номер", "Организация", "Головная организация", "Место проведения", "Дата начала проверки", "Дата окончания проверки", "Ответственный", "Область проверки"},
    }

    for rows.Next() {
        var ID, Name, HeadCompNm, ChArea, Responsible, AreaCheck string
        var StDate, EndDate time.Time

        rows.Scan(
            &ID,
            &Name,
            &HeadCompNm,
            &ChArea,
            &StDate,
            &EndDate,
            &Responsible,
            &AreaCheck,
        )
        data.Rows = append(data.Rows, []string{
            ID,
            Name,
            HeadCompNm,
            ChArea,
            StDate.Format("2025-01-01"),
            EndDate.Format("2025-01-01"),
            Responsible,
            AreaCheck,
        })
    }
    return data, nil
}

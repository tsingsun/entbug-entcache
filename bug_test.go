package bug

import (
	"ariga.io/entcache"
	"context"
	"entgo.io/bug/ent/securityjournal"
	"entgo.io/bug/ent/securityposition"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"net"
	"strconv"
	"testing"
	"time"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"entgo.io/bug/ent"
	"entgo.io/bug/ent/enttest"
)

func TestBugSQLite(t *testing.T) {
	drvori, err := sql.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		t.Error(err)
	}
	client := buildCacheClient(t, drvori)
	defer client.Close()
	test(t, client)
}

func TestBugMySQL(t *testing.T) {
	for version, port := range map[string]int{"56": 3306, "57": 3307, "8": 3308} {
		addr := net.JoinHostPort("localhost", strconv.Itoa(port))
		t.Run(version, func(t *testing.T) {
			drvori, err := sql.Open(dialect.MySQL, fmt.Sprintf("root:pass@tcp(%s)/test?parseTime=True", addr))
			if err != nil {
				t.Error(err)
				return
			}
			client := buildCacheClient(t, drvori)
			defer client.Close()
			test(t, client)
		})
	}
}

func TestBugPostgres(t *testing.T) {
	for version, port := range map[string]int{"10": 5430, "11": 5431, "12": 5432, "13": 5433, "14": 5434} {
		t.Run(version, func(t *testing.T) {
			drvori, err := sql.Open(dialect.Postgres, fmt.Sprintf("host=localhost port=%d user=postgres dbname=test password=pass sslmode=disable", port))
			if err != nil {
				t.Error(err)
				return
			}
			client := buildCacheClient(t, drvori)
			defer client.Close()
			test(t, client)
		})
	}
}

func TestBugMaria(t *testing.T) {
	for version, port := range map[string]int{"10.5": 4306, "10.2": 4307, "10.3": 4308} {
		t.Run(version, func(t *testing.T) {
			addr := net.JoinHostPort("localhost", strconv.Itoa(port))
			drvori, err := sql.Open(dialect.MySQL, fmt.Sprintf("root:pass@tcp(%s)/test?parseTime=True", addr))
			if err != nil {
				t.Error(err)
				return
			}
			client := buildCacheClient(t, drvori)
			defer client.Close()
			test(t, client)
		})
	}
}

func buildCacheClient(t *testing.T, drvori dialect.Driver) *ent.Client {
	drv := entcache.NewDriver(drvori, entcache.TTL(time.Minute*5))
	client := enttest.NewClient(t, enttest.WithOptions(ent.Driver(drv)))
	return client
}

func test(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	client.SecurityPosition.Delete().Where(securityposition.IDIn(8470, 8471)).ExecX(ctx)
	bulk := []*ent.SecurityPositionCreate{
		client.SecurityPosition.Create().SetID(8470).SetParentID(1).SetAccountID(3).SetSecurityAccountID(1).SetPositionType(1).
			SetCdDirection(-1).SetProjectID(33).SetProductID(51).SetMaterialID(123934).SetMaterialNo("688280.SH").SetMultiplier(1.0000).
			SetBalance(16945.0000).SetAvailable(16945.0000).SetFreeze(0.0000).SetAfloat(0.0000).SetUnit("1").SetCurrency("CNY").SetPrice(12.330000).
			SetAmount(202601.93).SetCostAmount(202601.93).SetCost(11.956443).SetFxRate(1.00000000).SetStlCurrency("CNY").SetStlAmount(202601.93).
			SetStlCost(11.956443).SetStlMargin(101300.97).SetStlCostAmount(202601.93).SetStlMarginLv(0.500000).SetStlValPrice(12.330000).SetCreatedAt(time.Now()).
			SetUniqueTag("0abe3765c08423a7f00521047036fe7a").SetOrgID(1000),
		client.SecurityPosition.Create().SetID(8471).SetParentID(1).SetAccountID(3).SetSecurityAccountID(1).SetPositionType(1).
			SetCdDirection(-1).SetProjectID(33).SetProductID(51).SetMaterialID(4107).SetMaterialNo("300472.SZ").SetMultiplier(1.0000).
			SetBalance(9000.0000).SetAvailable(9000.0000).SetFreeze(0.0000).SetAfloat(0.0000).SetUnit("1").SetCurrency("CNY").SetPrice(12.250000).
			SetAmount(103428.00).SetCostAmount(103428.00).SetCost(11.492000).SetFxRate(1.00000000).SetStlCurrency("CNY").SetStlAmount(103428.00).
			SetStlCost(11.492000).SetStlMargin(51714.00).SetStlCostAmount(103428.00).SetStlMarginLv(0.500000).SetStlValPrice(12.250000).SetCreatedAt(time.Now()).
			SetUniqueTag("640637ebdc0a64561e8b824b7f12e3fc").SetOrgID(1000),
	}
	client.SecurityPosition.CreateBulk(bulk...).SaveX(ctx)

	for i := 0; i < 2; i++ {
		datas, err := client.SecurityPosition.Query().Where(securityposition.AccountIDIn(3)).
			Order(ent.Asc(securityposition.FieldProjectID, securityposition.FieldProductID)).
			All(context.Background())
		if err != nil {
			t.FailNow()
		}
		if datas[0].ID == 8470 && *datas[0].MaterialNo != "688280.SH" {
			t.Error(i, *datas[0].MaterialNo)
		}
		if datas[1].ID == 8471 && *datas[1].MaterialNo != "300472.SZ" {
			t.Error(i, *datas[1].MaterialNo)
		}

		client.SecurityJournal.Query().
			Where(securityjournal.AccountIDIn(3),
				securityjournal.ChangeType("2"), securityjournal.IsDayBooking("Y"),
			).Order(ent.Desc(securityjournal.FieldID)).
			AllX(context.Background())
	}
}

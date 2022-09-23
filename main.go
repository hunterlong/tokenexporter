package main

import (
	"bufio"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"net/http"
	"os"
	"strings"
    "strconv"
	"time"
	"io/ioutil"
	"encoding/json"
	"sync"
	"sort"
)

var (
	allWatching []*Watching
	port        string
	updates     string
	eth         *ethclient.Client
)

type Watching struct {
	Name    string
	Address string
	Balance string
	Host    string
}

var tokenList []TokenList

//
// Connect to geth server
func ConnectionToGeth(url string) error {
	var err error
	eth, err = ethclient.Dial(url)
	return err
}

func tokenCaller(address common.Address) (*MainCaller, error) {
	caller, err := NewMainCaller(address, eth)
	if err != nil {
		return nil, err
	}
	return caller, err
}

//
// Fetch ETH balance from Geth server
func GetTokenBalance(token, address string, decimals int) string {
	caller, _ := tokenCaller(common.HexToAddress(token))
	balance, _ := caller.BalanceOf(nil, common.HexToAddress(address))
	corrected := ToDecimals(balance, decimals)
	return corrected
}


type Order struct {
	Index int
	Data string
}


func FindOrder(o []Order, index int) Order {
	var n Order
	for _, v := range o {
		if v.Index == index {
			n = v
		}
	}
	return n
}


func Resort(o []Order) []Order {
	order := []int{}
	for _, v := range o {
		order = append(order, v.Index)
	}
	sort.Ints(order)
	newOrder := []Order{}
	for _, v := range order {
		newOrder = append(newOrder, FindOrder(o, v))
	}
	return newOrder
}

func exp_func(x int, y int64) (value *big.Int) {

    exp:=strconv.FormatInt(y, 2)
    v := big.NewInt(int64(x))
    for i := 1; i < len(exp); i++ {
        v.Mul(v, v)
            if(exp[i]=='1') {
            v.Mul(v, big.NewInt(int64(x)))
        }
    }
    return v
}
//
// CONVERT USING THE DECIMALS
func ToDecimals(o *big.Int, decimals int) string{
    if o.String()=="0" {
        return "0"
    }
	pul, int := big.NewFloat(0), big.NewFloat(0)
	int.SetInt(o)
    bigDec := exp_func(10, int64(decimals))
    fDec := new(big.Float).SetInt(bigDec)
	pul = new(big.Float).Quo(int, fDec)
	fmt.Printf("pul  %f \n", pul)
    result := pul.String()
	return result
}

//
// HTTP response handler for /metrics
func MetricsHttp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, strings.Join(outData, "\n"))
}

//
// Open the addresses.txt file (name:address)
func OpenAddresses(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		object := strings.Split(scanner.Text(), ":")
		if common.IsHexAddress(object[1]) {
			w := &Watching{
				Name:    object[0],
				Address: object[1],
				Host: object[2],
			}
			allWatching = append(allWatching, w)
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return err
}


//
// Open the addresses.txt file (name:address)
func OpenTokenList(filename string) error {
	var jsonList []TokenList
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &jsonList)
	if err != nil {
		return err
	}
	tokenList = jsonList
	return err
}

var outData []string

func main() {
	gethUrl := os.Getenv("GETH")
	port = os.Getenv("PORT")

	err := OpenAddresses("addresses.txt")
	if err != nil {
		panic(err)
	}

	err = OpenTokenList("tokens-list.json")
	if err != nil {
		panic(err)
	}

	err = ConnectionToGeth(gethUrl)
	if err != nil {
		panic(err)
	}

	// check address balances
	go func() {
		for {
			pendingData := []Order{}
			pendingText := []string{}
			t1 := time.Now()
			guard := make(chan struct{}, 16)
			for _, v := range allWatching {
				wg := new(sync.WaitGroup)
				for i, tk := range tokenList {
					token := tk
					wg.Add(1)
					go func(i int, tk TokenList, v *Watching, wg *sync.WaitGroup) {
						guard <- struct{}{}
						balance := GetTokenBalance(tk.Address, v.Address, tk.Decimal)
						data := fmt.Sprintf("token_balance{host=\"%v\",name=\"%v\",symbol=\"%v\",contract=\"%v\",address=\"%v\"} %v", v.Host, v.Name, tk.Symbol, tk.Address, v.Address, balance)
						order := Order{i, data}
						pendingData = append(pendingData, order)
						<-guard
						wg.Done()
					}(i, token, v, wg)
				}
				wg.Wait()
				pendingData = Resort(pendingData)
				watchingText := OrderToString(pendingData)
				pendingText = append(pendingText, watchingText)
				pendingData = []Order{}
			}
			t2 := time.Now()
			diff := t2.Sub(t1)
			data := fmt.Sprintf("token_query_seconds{name=\"%v\"} %v", "all", diff.Seconds())
			pendingText = append(pendingText, data)
			outData = pendingText
			fmt.Printf("Token balance queries completed in %f seconds\n", diff.Seconds())
			time.Sleep(600 * time.Second)
		}
	}()

	fmt.Printf("ETHexporter has started on port %v using Geth server: %v\n", port, gethUrl)
	http.HandleFunc("/metrics", MetricsHttp)
	panic(http.ListenAndServe("0.0.0.0:"+port, nil))
}


func OrderToString(orders []Order) string {
	var outgoing []string
	for _, v := range orders {
		outgoing = append(outgoing, v.Data)
	}
	return strings.Join(outgoing, "\n")
}


type Tokens struct {
	List []TokenList
}


type TokenList struct {
	Address string `json:"address"`
	Symbol  string `json:"symbol"`
	Decimal int    `json:"decimal"`
	Type    string `json:"type"`
}

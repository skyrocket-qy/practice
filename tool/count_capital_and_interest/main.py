MIN_BUY_UNITS = 1000

class DIVIDEND_INTERVAL:
    YEAR = 1
    HALF_YEAR = 2
    QUARTER = 3
    MONTH = 12

rate = 0.045
year = 28
month_deposit = 40000
dividend_interval = DIVIDEND_INTERVAL.HALF_YEAR

def main():
    money = bond = 0
    dividend_rate = rate / int(dividend_interval)
    n = year * int(dividend_interval)
    deposit = month_deposit * 12 / int(dividend_interval)

    # suppose start from XXXX-01-01
    for _ in range(n):
        interest = bond * dividend_rate
        money += interest + deposit
        
        can_buy = money - (money % MIN_BUY_UNITS)
        bond += can_buy
        money -= can_buy

    print(f"money: {money}, bond: {bond}")

if __name__ == '__main__':
    main()
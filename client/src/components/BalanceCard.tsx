import { Card } from "./Card";


export const BalanceCard = ({
    date,
    asset,
    de,
    debt,
}: BalanceCardInterface) => {
    return(
        <Card>
            <div className="flex justify-between items-center px-4 py-2">
                <p className="text-pf-gray-100 font-bold text-3xl">Balance</p>
                <p className="text-pf-gray-100 font-bold text-2xl">{date}</p>
            </div>
            <div className="mt-5 flex justify-between items-center px-4 py-2">
                <p className="text-pf-gray-100 font-bold text-2xl">Asset</p>
                <p className="text-pf-gray-100 font-bold text-2xl">D/E</p>
            </div>
            <div className="flex justify-between items-center px-4 py-2">
                <p className="text-pf-accent-2 font-bold text-2xl">{asset.toLocaleString(undefined, { maximumFractionDigits: 2 })}</p>
                <p className="text-pf-accent-2 font-bold text-2xl">{de.toLocaleString(undefined, { maximumFractionDigits: 2 })}</p>
            </div>
            <div className="mt-5 flex justify-between items-center px-4 py-2">
                <p className="text-pf-gray-100 font-bold text-2xl">Debt</p>
            </div>
            <div className="flex justify-between items-center px-4 py-2">
                <p className="text-pf-accent-2 font-bold text-2xl">{debt.toLocaleString(undefined, { maximumFractionDigits: 2 })}</p>
            </div>
        </Card>
    );
}
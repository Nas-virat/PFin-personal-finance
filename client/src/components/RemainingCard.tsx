import { Card } from "./Card";


export const RemainingCard = ({
    date,
    revenue,
    expense,
    remaining,
    credit
}: RemainingInterface) => {
    return (
        <Card>
            <div className="flex justify-between items-center px-4 py-2">
                <p className="text-pf-gray-100 font-bold text-3xl">Remaining</p>
                <p className="text-pf-gray-100 font-bold text-2xl">{date}</p>
            </div>
            <div className="flex justify-between items-center px-4 py-2">
                <p className="text-pf-accent-2 font-bold text-2xl">{remaining.toLocaleString(undefined, { maximumFractionDigits: 2 })}</p>
            </div>
            <div className="mt-5 flex justify-between items-center px-4 py-2">
                <p className="text-pf-gray-100 font-bold text-2xl">Expense</p>
                <p className="text-pf-gray-100 font-bold text-2xl">Revenue</p>
                <p className="text-pf-gray-100 font-bold text-2xl">Credit</p>
            </div>
            <div className="flex justify-between items-center px-4 py-2">
                <p className="text-pf-accent-2 font-bold text-2xl">{revenue.toLocaleString(undefined, { maximumFractionDigits: 2 })}</p>
                <p className="text-pf-accent-2 font-bold text-2xl">{expense.toLocaleString(undefined, { maximumFractionDigits: 2 })}</p>
                <p className="text-pf-accent-2 font-bold text-2xl">{credit.toLocaleString(undefined, { maximumFractionDigits: 2 })}</p>
            </div>
        </Card>
    );
}
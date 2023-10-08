

export const RemainingCard = ({
    date,
    revenue,
    expense,
    remaining,
    credit
}: RemainingInterface) => {
    return (
        <div className="w-5/6 mx-9 bg-pf-gray-900 rounded-xl">
            <div className="flex justify-between items-center px-4 py-2">
                <p className="text-pf-gray-100 font-bold text-xl">Remaining</p>
                <p className="text-pf-gray-100 font-bold text-xl">{date}</p>
            </div>
            <div className="flex justify-between items-center px-4 py-2">
                <p className="text-pf-accent-2 font-bold text-xl">{remaining.toLocaleString(undefined, { maximumFractionDigits: 2 })}</p>
            </div>
            <div className="mt-5 flex justify-between items-center px-4 py-2">
                <p className="text-pf-gray-100 font-bold text-xl">Expense</p>
                <p className="text-pf-gray-100 font-bold text-xl">Revenue</p>
                <p className="text-pf-gray-100 font-bold text-xl">Credit</p>
            </div>
            <div className="flex justify-between items-center px-4 py-2">
                <p className="text-pf-accent-2 font-bold text-xl">{revenue.toLocaleString(undefined, { maximumFractionDigits: 2 })}</p>
                <p className="text-pf-accent-2 font-bold text-xl">{expense.toLocaleString(undefined, { maximumFractionDigits: 2 })}</p>
                <p className="text-pf-accent-2 font-bold text-xl">{credit.toLocaleString(undefined, { maximumFractionDigits: 2 })}</p>
            </div>
        </div>
    );
}
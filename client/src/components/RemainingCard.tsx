import { Card } from "./Card";
import { Content, SubHeader } from "./SubHeader";


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
                <SubHeader text={date}/>
            </div>
            <div className="flex justify-between items-center px-4 py-2">
                <p className="text-pf-accent-2 font-bold text-2xl">{remaining.toLocaleString(undefined, { maximumFractionDigits: 2 })}</p>
            </div>
            <div className="mt-5 flex justify-between items-center px-4 py-2">
                <SubHeader text="Revenue"/>
                <SubHeader text="Expense"/>
                <SubHeader text="Credit"/>
            </div>
            <div className="flex justify-between items-center px-4 py-2">
                <Content text={revenue.toLocaleString(undefined, { maximumFractionDigits: 2 })}/>
                <Content text={expense.toLocaleString(undefined, { maximumFractionDigits: 2 })}/>
                <Content text={credit.toLocaleString(undefined, { maximumFractionDigits: 2 })}/>
            </div>
        </Card>
    );
}
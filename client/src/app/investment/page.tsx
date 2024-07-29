"use client";
import React, { useState } from "react";

import { AddButton } from "@/components/Addbutton";
import { useRouter } from "next/navigation";
import dayjs, { Dayjs } from "dayjs";
import { Card } from "@/components/Card";
import { DoughnutChart } from "@/components/chart/DoughnutChart";
import { HeaderCard } from "@/components/HeaderCard";
import { PageContainer } from "@/components/PageContainer";
import { TableInfo } from "@/components/TableInfo";

export default function Page() {
  const router = useRouter();
  const [date, setDate] = useState<Dayjs>(dayjs());

  return (
    <div>
      <div className="flex justify-between mx-32 text-pf-gray-900 font-bold text-2xl">
        <h1 className="text-[48px]">Investment</h1>
        <AddButton
          text="Add Transaction"
          func={() =>
            router.push(
              "/transaction?date=" +
                date.date() +
                "&month=" +
                (date.month() + 1) +
                "&year=" +
                date.year() +
                "",
            )
          }
        />
      </div>
      <PageContainer>
        <div className="w-full flex justify-center">
          <Card>
            <HeaderCard text="Total Investment" />
            <TableInfo
              data={[1, 2, 3]}
              columns={["Asset Type", "Amount"]}
              total={30}
            />
          </Card>
          <Card>
            <div className="flex justify-center">
              <DoughnutChart
                data={[10, 20, 30]}
                labels={["Equity", "Debt", "Gold"]}
                backgroundColor={["#FF6384", "#36A2EB", "#FFCE56"]}
              />
            </div>
          </Card>
        </div>
      </PageContainer>
      <PageContainer>
        <div className="w-full flex justify-center">
          <Card>
            <HeaderCard text="Total Stock" />
            <TableInfo
              data={[
                { 1: 2, 3: 4 },
                { 3: 4, 4: 5 },
                { 5: 6, 6: 7 },
              ]}
              columns={["Stock Name", "Amount"]}
              total={30}
            />
          </Card>
          <Card>
            <div className="flex justify-center">
              <DoughnutChart
                data={[10, 20, 30]}
                labels={["Equity", "Debt", "Gold"]}
                backgroundColor={["#FF6384", "#36A2EB", "#FFCE56"]}
              />
            </div>
          </Card>
        </div>
      </PageContainer>
      <PageContainer>
        <div className="w-full flex justify-center">
          <Card>
            <HeaderCard text="Total Fund" />
            <TableInfo
              data={[1, 2, 3]}
              columns={["Asset Type", "Amount"]}
              total={30}
            />
          </Card>
          <Card>
            <div className="flex justify-center">
              <DoughnutChart
                data={[10, 20, 30]}
                labels={["Equity", "Debt", "Gold"]}
                backgroundColor={["#FF6384", "#36A2EB", "#FFCE56"]}
              />
            </div>
          </Card>
        </div>
      </PageContainer>
      <PageContainer>
        <div className="w-full flex justify-center">
          <Card>
            <HeaderCard text="Total Bond" />
            <TableInfo
              data={[1, 2, 3]}
              columns={["Asset Type", "Amount"]}
              total={30}
            />
          </Card>
          <Card>
            <div className="flex justify-center">
              <DoughnutChart
                data={[10, 20, 30]}
                labels={["Equity", "Debt", "Gold"]}
                backgroundColor={["#FF6384", "#36A2EB", "#FFCE56"]}
              />
            </div>
          </Card>
        </div>
      </PageContainer>
    </div>
  );
}

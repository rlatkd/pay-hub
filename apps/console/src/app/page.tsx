import { Sidebar } from '@/widgets/layout/ui/Sidebar';
import { StatCard } from '@/shared/ui/StatCard';
import { BillingWidget } from '@/widgets/dashboard/ui/BillingWidget';
import { Users, FileText, DollarSign, TrendingUp } from 'lucide-react';

export default function DashboardPage() {
  return (
    <div className="flex min-h-screen bg-[#F4F7FE]">
      <Sidebar />
      <main className="flex-1 lg:ml-[290px] p-8 md:p-12 transition-all">
        <header className="flex flex-col md:flex-row justify-between items-start md:items-center mb-10 gap-4">
          <div>
            <h1 className="text-[34px] font-bold text-[#2B3674] tracking-tight leading-tight">TODO월 청구 현황</h1>
          </div>
          <div className="flex items-center gap-4 bg-white p-3 pl-5 rounded-full shadow-sm shadow-gray-100">
             <input type="text" placeholder="검색..." className="bg-[#F4F7FE] rounded-full px-4 py-2 text-sm focus:outline-none w-40 hidden md:block" />
             <span className="text-sm font-bold text-[#2B3674] px-3 border-l">김상훈님</span>
             <div className="w-10 h-10 bg-blue-100 rounded-full bg-[url('https://i.pravatar.cc/100')] bg-cover"></div>
          </div>
        </header>
        <section className="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-4 gap-6 mb-10">
          <StatCard 
            label="회원 현황" value="251명" subText="신규 157 | 활성 251"
            icon={Users} iconColor="text-[#4318FF]" iconBg="bg-[#F4F7FE]"
          />
          <StatCard 
            label="계약 현황" value="204건" subText="신규 14 | 완료예정 49"
            icon={FileText} iconColor="text-[#4318FF]" iconBg="bg-[#F4F7FE]"
          />
          <StatCard 
            label="청구 현황" value="₩3.7억" subText="완납 3.2억 | 미납 251만"
            icon={DollarSign} iconColor="text-[#05CD99]" iconBg="bg-[#E6F9F3]"
          />
          <StatCard 
            label="전월 대비 매출" value="-16.7%" subText="회원 +188.9% 증가"
            icon={TrendingUp} iconColor="text-[#EE5D50]" iconBg="bg-[#FEEFEF]"
          />
        </section>
        <section className="grid grid-cols-1 xl:grid-cols-5 gap-8 h-auto xl:h-[600px]">
          <div className="xl:col-span-3 bg-white rounded-[30px] p-8 shadow-sm flex flex-col">
            <div className="flex justify-between items-center mb-8">
              <h3 className="text-2xl font-bold text-[#2B3674]">2024년 8월</h3>
              <div className="flex gap-3">
                 <button className="bg-[#F4F7FE] text-[#4318FF] px-4 py-2 rounded-xl text-sm font-bold hover:bg-blue-100 transition">이전</button>
                 <button className="bg-[#F4F7FE] text-[#4318FF] px-4 py-2 rounded-xl text-sm font-bold hover:bg-blue-100 transition">다음</button>
              </div>
            </div>
            <div className="flex-1 border-2 border-dashed border-gray-200 rounded-3xl bg-[#F8F9FA] flex items-center justify-center text-[#A3AED0] font-medium text-lg">
               캘린더 컴포넌트 자리
            </div>
          </div>
          <div className="xl:col-span-2">
            <BillingWidget />
          </div>
        </section>
      </main>
    </div>
  );
}

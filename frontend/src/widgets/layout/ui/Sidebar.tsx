import { Home, Users, FileText, CreditCard, Box, BarChart2, Settings } from 'lucide-react';

export const Sidebar = () => {
  const menuItems = [
    { name: '대시보드', icon: Home, active: true },
    { name: '회원', icon: Users, active: false },
    { name: '계약', icon: FileText, active: false },
    { name: '청구', icon: CreditCard, active: false },
    { name: '상품', icon: Box, active: false },
    { name: '통계', icon: BarChart2, active: false },
  ];

  return (
    <aside className="w-[290px] bg-white h-screen fixed left-0 top-0 border-r border-gray-100 hidden lg:flex flex-col p-6 z-50 font-medium">
      <div className="flex items-center gap-2 mb-12 px-4 py-4">
        <div className="text-3xl font-black text-[#2B3674] tracking-tight">
          PAYHUB
        </div>
      </div>
      <nav className="flex flex-col gap-3 flex-1">
        {menuItems.map((item) => (
          <div
            key={item.name}
            className={`flex items-center gap-4 px-5 py-4 rounded-2xl cursor-pointer transition-all duration-200 ${
              item.active 
                ? 'bg-[#4318FF] text-white font-bold shadow-lg shadow-blue-500/30' 
                : 'text-[#A3AED0] hover:bg-gray-100 hover:text-[#2B3674]'
            }`}
          >
            <item.icon size={22} />
            <span className="text-base">{item.name}</span>
          </div>
        ))}
      </nav>
       <div className="mt-auto mb-6 px-5 py-4 text-[#A3AED0] hover:text-[#2B3674] flex items-center gap-4 cursor-pointer font-medium">
          <Settings size={22} />
          <span>설정</span>
       </div>
    </aside>
  );
};

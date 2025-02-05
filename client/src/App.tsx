import React, { useEffect, useState } from "react";
import { Table } from "antd";
import { format } from "date-fns";
import { ru } from "date-fns/locale";

interface ContainerStatus {
  id: number;
  ip_address: string;
  status: string;
  last_checked: string;
}

const API_URL = "/api/containers";

const formatDate = (isoString: string) => {
  return format(new Date(isoString), "dd.MM.yyyy HH:mm:ss", { locale: ru });
};

const App: React.FC = () => {
  const [data, setData] = useState<ContainerStatus[]>([]);

  const fetchData = async () => {
    try {
      const response = await fetch(API_URL);
      const result = await response.json();
      setData(result);
    } catch (error) {
      console.error("Ошибка загрузки данных:", error);
    }
  };

  useEffect(() => {
    fetchData();
    const interval = setInterval(fetchData, 5000);
    return () => clearInterval(interval);
  }, []);

  const columns = [
    { title: "IP-адрес контейнера", dataIndex: "ip_address", key: "ip_address" },
    { title: "Статус контейнера", dataIndex: "status", key: "status" },
    {
      title: "Последняя проверка контейнера",
      dataIndex: "last_checked",
      key: "last_checked",
      render: (text: string) => formatDate(text),
    },
  ];

  return (
    <div style={{ padding: 20 }}>
      <h1>Мониторинг контейнеров</h1>
      <Table dataSource={data} columns={columns} rowKey="id" />
    </div>
  );
};

export default App;

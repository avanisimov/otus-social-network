import http from 'k6/http';
import { sleep, check } from 'k6';
import { Trend } from 'k6/metrics';

// Метрики для /user/get
export let getUserLatency = new Trend('latency_get_user');

// Метрики для /user/search
export let searchUserLatency = new Trend('latency_search_user');

export let options = {
    scenarios: {
        get_user: {
        executor: 'constant-arrival-rate',
        exec: 'getUserScenario',
        rate: 100,                // начальный RPS
        timeUnit: '1s',
        duration: '1m',
        preAllocatedVUs: 50,      // заранее созданные VU
        maxVUs: 500,              // максимум, сколько может создать k6
        },

        search_user: {
        executor: 'constant-arrival-rate',
        exec: 'searchUserScenario',
        rate: 30,                 // RPS для поиска
        timeUnit: '1s',
        duration: '1m',
        preAllocatedVUs: 50,
        maxVUs: 500,
        },
    },

    thresholds: {
        // Важные метрики для отчёта
        'latency_get_user': ['p(95)<50'],        // например: p95 < 50мс
        'latency_search_user': ['p(95)<200'],    // может быть большой до индекса
        'http_req_failed': ['rate<0.01'],        // <1% ошибок
    },
};

const userIds = [
    '71b83d6f-be08-4ee9-b0c5-c2d631b6abe1',
    'b736e0cb-b490-4c79-82e7-24b14cbe4d07',
    '717c64cc-832b-4f7a-82bb-a03672da0bc6',
    '1748bc2e-a015-4e14-99ac-5796f882a805',
    '46236a1e-c020-446b-9876-7c1bb7d1cae7',
    '6f512539-0e2d-4a70-b861-6063c2fe6395',
    '7e33c10b-b7a2-4543-9beb-3a47137ac215',
    'c4d0c9e6-1647-4b4b-9f53-6c2f14d96cda',
    '80d7dd32-4b4e-4200-827e-5f6aeb75cc7f',
    '4a9de2d6-a308-4bf9-a713-84c496c3a93c',
    '1a9cf0a1-2eb4-413a-bb49-9d86109a7852',
    '1dfc1d82-487c-48eb-85fe-1566b36457b8',
    'b9b23125-4171-45b8-9f56-4a6d69a0ee99',
    '276fa806-9ca0-4b72-8b61-da7152ff859c',
    '16723152-8e34-4385-82df-c0ebf77010a1',
    '45543d04-4bc8-4175-9ed2-f67af10331cc',
    'ed7564ab-e409-4864-ac76-9985be7c5cd5',
    '959c55d3-7924-48f2-b4e1-7f5d67a9f6e0',
    'f8cee3b4-b5ae-43f1-9684-e11c38e064ac',
    '0226d4af-2bec-4b56-9c2f-5322fa63e208',
    '6ea8425d-5a41-4974-a565-c86762a3e8cd',
    'cae0ac31-3e81-4678-ad61-bc7653f0b816',
    '5333df35-f172-4721-93f4-18b64358f770',
    '76ff73eb-1c0b-4f7a-b4a9-aeb924d30276',
    'b3d3eb82-c032-4e41-ad2f-62ffaf92accc',
    '680ddb1f-77be-4558-8ee5-b477b997bf53',
    '8b558af1-9197-46a0-9aac-ca01753ff41f',
    '3b4ea5ef-85ae-408b-aa23-263ffa496616',
    '1b86f193-48c6-4f06-b107-d3ef50b45ae6',
    '5f3ba376-643a-4c0b-94ab-6cf5da089230',
    '8608f9ad-e57e-41c7-821d-5a310a96acd8',
    'd1c6c86c-cf83-4cdf-af4a-27f62cdb72a5',
    'f697513a-935b-4caf-8e3f-7697f2c4da4b',
    '88e8174f-6f13-4c7b-8225-af2798c02d38',
    '766b9e77-7e9d-4742-a49b-3a58d9ea0849',
    '51272dc3-46c0-49cb-89fa-513198d44161',
    'e7681759-6472-4cb5-9808-f9109c5a7f8b',
    'db01f243-c116-4843-be94-77195cc1029b',
    'f6b73cdb-dbe2-4a2b-8bb5-94f896014800',
    '19f32e60-9db2-4deb-bdaa-2c4b7b1d4993',
    '8345fb3d-48cc-468f-889c-c8832d5e0bc3',
    'd654a763-493e-4bd3-b1f4-1da7823834a5',
    '9db04cfc-b188-4b59-bbe5-f31a11fb24c9',
    'cd9779cc-135d-4568-ae16-9c3b1477aca2',
    'e4217a3a-2fb4-4c69-b0c6-c007af911567',
    '4a64d591-9912-470a-8dc7-fb43dcd0c6bb',
    '8d49a299-5a22-489c-a1e5-c72f94361128',
    'affd55cc-9b39-41c8-a818-eb9a6be5cd55',
    '55df1df1-43af-4d27-b893-35d762c8bacb',
    'a3ab0f5b-58f5-4c31-be9e-0c2456a2814a'
];

function getUserId() {
    const chance = Math.random();

    // 90% — реальные пользователи
    if (chance < 0.9) {
        return randomItem(userIds);
    }

    // 10% — случайные (не существующие)
    return randomUUID();
}

function randomItem(arr) {
    return arr[Math.floor(Math.random() * arr.length)];
}
function randomUUID() {
    // простой фейковый UUID (для несуществующих пользователей)
    return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
        const r = Math.random() * 16 | 0;
        const v = c === 'x' ? r : (r & 0x3 | 0x8);
        return v.toString(16);
    });
}

// Генератор набора имён для поиска
const firstNames = ['Яро','Ива','Ром','Ерем','Твер','Эрн','Рома','Клав','Зино','Але','Радо','Вер','Луч','Сав','Анн','Руб','Аве','Натан','Афин','Касьян','Зин','Вени','Апо','Авдей', 'Фео', 'Баже'];
const lastNames = ['Каз','Ков','Саф','Дем','Каб','Сав','Тре','Ден','Мар','Кон','Лаз','Гри','Мур','Сим','Лав','Жук','Рож','Орл','Пах','Зык'];

function randomFrom(arr) {
    return arr[Math.floor(Math.random() * arr.length)];
}

// --------------------------
// Сценарий 1: /user/get/{id}
// --------------------------
export function getUserScenario() {
    const id = getUserId();
    const url = `http://localhost:8080/user/get/${id}`;

    const res = http.get(url);

    getUserLatency.add(res.timings.duration);

    check(res, {
        'status is 200': (r) => r.status === 200,
    });

    sleep(0.2);
}

// --------------------------
// Сценарий 2: /user/search
// --------------------------
export function searchUserScenario() {
    const fn = randomFrom(firstNames);
    const ln = randomFrom(lastNames);

    const url = `http://localhost:8080/user/search?firstName=${fn}&secondName=${ln}&limit=20`;

    const res = http.get(url);

    searchUserLatency.add(res.timings.duration);

    check(res, {
        'status is 200': (r) => r.status === 200,
    });

    sleep(0.2);
}
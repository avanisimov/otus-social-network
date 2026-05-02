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
    "5fead6db-b92d-4bf6-ac61-25bd2e6a8ad9",
"877416ef-4ec8-4a9e-b845-94fe333ed4ac",
"f18d0bb3-2f4f-4ca8-8eb1-3247bc967820",
"c9f30ac0-4705-4551-b4da-1022fe4b3416",
"4f2b50c2-0be5-4b87-87ca-c023e49c7cac",
"9975b561-5075-4407-a757-11d44e5abcf4",
"6a3489d1-0523-47c3-9869-68c7229b193d",
"6eb7b5a5-e5fb-4574-95ac-084be613c41c",
"f01b28cb-8cdb-44af-aa73-3a0183d66153",
"b9493060-bbf6-4159-b6b6-e0220b338f33",
"edf1427b-e171-407a-90c0-b8595eb75b1c",
"f706286c-e79f-45f8-9fe2-6117aa592c5a",
"ce2945be-112a-4ad6-ab06-e9bd616d85a8",
"d20b45cf-418c-4ae8-931f-465158480a7d",
"4996177b-65f2-4e6b-8215-e2fb933ac922",
"02a1210e-2834-4137-9cbb-1c20bb59f50e",
"624b3b1d-58bd-47aa-9606-4de6f7e5d3d0",
"6c8d499e-1ae0-4a34-b5db-377cb6a57c7d",
"97b9727a-6504-423c-97b7-474f21dfa18d",
"626fa2bd-478b-43ba-96b1-c09b4c6206ae"
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
const firstNames = ['Ост','Лав','Анг','Тер','Ост','Вик','Фил','Ила','Эле','Эрн','Нин','Адр','Агг','Мар','Евд','Спи','Сид','Рад','Пор','Вяч'];
const lastNames = ['Вес','Мат','Фро','Сми','Ефр','Хар','Сим','Гав','Боб','Лук','Фро','Фед','Леб','Гор','Бел','Ром','Кул','Бог','Сор','Кул'];

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
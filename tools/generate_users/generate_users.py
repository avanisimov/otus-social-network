import csv
import random
import hashlib
import argparse
from faker import Faker

fake = Faker("ru_RU")

hobbies = [
    "люблю путешествовать",
    "читаю научную фантастику",
    "играю в видеоигры",
    "занимаюсь спортом",
    "готовлю новые блюда",
    "слушаю музыку",
    "гуляю с собакой",
    "рисую",
    "катаюсь на велосипеде",
    "вяжу свитера",
    "хожу на рыбалку",
]

cities = [
    "Москва", "Санкт-Петербург", "Новосибирск", "Екатеринбург", "Казань",
    "Нижний Новгород", "Самара", "Омск", "Ростов-на-Дону", "Краснодар",
]


def random_bio():
    return f"{random.choice(hobbies)}, {random.choice(hobbies)}"


def random_birthdate():
    return fake.date_between(start_date="-70y", end_date="-5y").strftime("%Y-%m-%d")


def generate(count: int, output: str):
    with open(output, "w", newline="", encoding="utf-8") as f:
        writer = csv.writer(f)

        for i in range(count):
            first_name = fake.first_name()
            last_name = fake.last_name()
            birthdate = random_birthdate()
            bio = random_bio()
            city = random.choice(cities)

            password_raw = f"password123{i}".encode("utf-8")
            password_hash = hashlib.md5(password_raw).hexdigest()

            writer.writerow([
                first_name,
                last_name,
                birthdate,
                bio,
                city,
                password_hash,
            ])

            if i % 100000 == 0 and i > 0:
                print(f"[+] {i} users generated...")

    print(f"\n[✓] Done! Generated {count} users → {output}")


def main():
    parser = argparse.ArgumentParser(description="User CSV generator for highload social network")

    parser.add_argument(
        "--count",
        type=int,
        default=1_000_000,
        help="Number of users to generate (default: 1,000,000)"
    )

    parser.add_argument(
        "--output",
        type=str,
        default="users.csv",
        help="Output CSV file name (default: users.csv)"
    )

    args = parser.parse_args()

    print(f"[*] Generating {args.count:,} users into {args.output}...")
    generate(args.count, args.output)


if __name__ == "__main__":
    main()
import csv


def main() -> None:
    with open("expected.csv") as f:
        reader = csv.reader(f)

        expected = [row[0] for row in reader]

    with open("result.csv") as f:
        reader = csv.reader(f)

        result = [row[0] for row in reader]

    assert len(expected) == len(result), "Length missmatch"

    for row in expected:
        assert row in result, "Row missing"

    print("Test passed!")


if __name__ == "__main__":
    main()

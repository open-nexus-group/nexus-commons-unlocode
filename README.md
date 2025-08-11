# Readme

A package containing the entire UNLOCODE Dataset.
Everything is conveniently packaged into structs and collections with fast access through maps that are created when using the synthetic constructor.

You can theoretically access the data directly through their Slice Representation, but it's easier to use the data retrievel methods provided by the collections, which utilizes maps for the lookups instead of searching through the slice.

## Updating
To update to the latest unlocode dataset:

1. Update the type definitions if they have changed
2. Run the generator to generate the datasets

### Running the generator
go run ./generator/generator.go --type=unlocode --files=./generator/2024-2-UNLOC-1.csv,./generator/2024-2-UNLOC-2.csv,./generator/2024-2-UNLOC-3.csv


# To-Dos

- [] Calculate actual lat long values from the values that UNECE provides

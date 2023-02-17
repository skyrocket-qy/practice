<script context="module">
    export const columns = [
        { name: 'id'},
        { name: 'name' },
        { name: 'age' },
        { name: 'is_manager', format: (value) => (value ? '✔️' : '✖️') },
        { name: 'start_date' },
    ]

    export const rows = [
        { id: 1, name: 'Liz Lemon', age: 36, is_manager: true, start_date: '1999-02-28' },
        { id: 2, name: 'Jack Donaghy', age: 40, is_manager: true, start_date: '1997-03-05' },
        { id: 3, name: 'Tracy Morgan', age: 39, is_manager: false, start_date: '2002-07-12' },
        { id: 4, name: 'Jenna Maroney', age: 40, is_manager: false, start_date: '1999-02-28' },
        { id: 5, name: 'Kenneth Parcell', age: Infinity, is_manager: false, start_date: '1970-01-01' },
        { id: 6, name: 'Pete Hornberger', age: 42, is_manager: true, start_date: '2000-04-01' },
        { id: 7, name: 'Frank Rossitano', age: 36, is_manager: false, start_date: '2004-06-09' },
        { id: 8, name: null, age: null, is_manager: null, start_date: null },
    ]
</script>